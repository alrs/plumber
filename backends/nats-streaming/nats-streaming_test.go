package nats_streaming

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/batchcorp/plumber-schemas/build/go/protos/args"
	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"

	"github.com/batchcorp/plumber/backends/nats-streaming/stanfakes"
	"github.com/batchcorp/plumber/types"
	"github.com/batchcorp/plumber/validate"
)

var _ = Describe("Nats Streaming Backend", func() {
	var connOpts *opts.ConnectionOptions

	BeforeEach(func() {
		connOpts = &opts.ConnectionOptions{
			Conn: &opts.ConnectionOptions_NatsStreaming{
				NatsStreaming: &args.NatsStreamingConn{
					Dsn:             "localhost",
					UserCredentials: nil,
					ClusterId:       "test",
					ClientId:        "plumber",
					TlsOptions: &args.NatsStreamingTLSOptions{
						TlsSkipVerify: false,
						TlsCaCert:     []byte(`../../test-assets/ssl/ca.crt`),
						TlsClientCert: []byte(`../../test-assets/ssl/client.crt`),
						TlsClientKey:  []byte(`../../test-assets/ssl/client.key`),
					},
				},
			},
		}
	})

	Context("Name", func() {
		It("returns name", func() {
			Expect((&NatsStreaming{}).Name()).To(Equal(BackendName))
		})
	})

	Context("Close", func() {
		It("calls close on client", func() {
			fakeStan := &stanfakes.FakeConn{}
			fakeStan.CloseStub = func() error {
				return nil
			}
			err := (&NatsStreaming{stanClient: fakeStan}).Close(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(fakeStan.CloseCallCount()).To(Equal(1))
		})
	})

	Context("Test", func() {
		It("returns not implemented err", func() {
			err := (&NatsStreaming{}).Test(nil)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(types.NotImplementedErr))
		})
	})

	Context("validateBaseConnOpts", func() {
		It("validates conn presence", func() {
			err := validateBaseConnOpts(nil)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(validate.ErrMissingConnOpts))
		})
		It("validates conn config", func() {
			connOpts = &opts.ConnectionOptions{}
			err := validateBaseConnOpts(connOpts)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(validate.ErrMissingConnCfg))
		})
		It("validates MQTT presence", func() {
			connOpts = &opts.ConnectionOptions{
				Conn: &opts.ConnectionOptions_NatsStreaming{
					NatsStreaming: nil,
				},
			}
			err := validateBaseConnOpts(connOpts)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(validate.ErrMissingConnArgs))
		})
		It("validates TLS key", func() {
			args := connOpts.GetNatsStreaming()
			args.Dsn = "tls://localhost"
			args.TlsOptions.TlsClientKey = nil
			err := validateBaseConnOpts(connOpts)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ErrMissingTLSKey))
		})
		It("validates TLS Cert", func() {
			args := connOpts.GetNatsStreaming()
			args.Dsn = "tls://localhost"
			args.TlsOptions.TlsClientCert = nil
			err := validateBaseConnOpts(connOpts)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ErrMissingTlsCert))
		})
		It("validates TLS Certificate Authority", func() {
			args := connOpts.GetNatsStreaming()
			args.Dsn = "tls://localhost"
			args.TlsOptions.TlsCaCert = nil
			err := validateBaseConnOpts(connOpts)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ErrMissingTLSCA))
		})
		It("passes validation", func() {
			err := validateBaseConnOpts(connOpts)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
