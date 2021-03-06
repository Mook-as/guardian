package dns_test

import (
	"io/ioutil"
	"net"
	"os"

	"github.com/cloudfoundry-incubator/guardian/kawasaki/dns"
	"github.com/pivotal-golang/lager"
	"github.com/pivotal-golang/lager/lagertest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResolvFileCompiler", func() {
	var (
		hostResolvConfPath string

		log      lager.Logger
		compiler *dns.ResolvFileCompiler
	)

	BeforeEach(func() {
		log = lagertest.NewTestLogger("test")
	})

	JustBeforeEach(func() {
		compiler = &dns.ResolvFileCompiler{
			HostResolvConfPath: hostResolvConfPath,
			HostIP:             net.ParseIP("254.253.252.251"),
		}
	})

	writeFile := func(filePath, contents string) {
		f, err := os.Create(filePath)
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		_, err = f.Write([]byte(contents))
		Expect(err).NotTo(HaveOccurred())
	}

	Context("when the host resolv.conf file does not exist", func() {
		BeforeEach(func() {
			hostResolvConfPath = "/does/not/exist.conf"
		})

		It("should return an error", func() {
			_, err := compiler.Compile(log)
			Expect(err).To(MatchError(ContainSubstring(("reading file '/does/not/exist.conf'"))))
		})
	})

	Context("when the host resolv.conf exists", func() {
		BeforeEach(func() {
			f, err := ioutil.TempFile("", "")
			Expect(err).NotTo(HaveOccurred())
			hostResolvConfPath = f.Name()
		})

		AfterEach(func() {
			Expect(os.Remove(hostResolvConfPath)).To(Succeed())
		})

		Context("and the host is running DNS", func() {
			BeforeEach(func() {
				writeFile(hostResolvConfPath, "nameserver 127.0.0.1\n")
			})

			It("should make the container use host DNS", func() {
				contents, err := compiler.Compile(log)
				Expect(err).NotTo(HaveOccurred())

				Expect(string(contents)).To(Equal("nameserver 254.253.252.251\n"))
			})
		})

		Context("and the host is not running DNS", func() {
			var resolvConfContents string

			BeforeEach(func() {
				resolvConfContents = "nameserver 127.0.0.1\nnameserver 8.8.4.4\n"
				writeFile(hostResolvConfPath, resolvConfContents)
			})

			It("should copy the host's resolv.conf", func() {
				contents, err := compiler.Compile(log)
				Expect(err).NotTo(HaveOccurred())

				Expect(string(contents)).To(Equal(resolvConfContents))
			})
		})
	})
})
