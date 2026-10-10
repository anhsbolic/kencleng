package probe
import (
 "context"
 "github.com/coreos/go-oidc/v3/oidc"
 "golang.org/x/oauth2"
)
func Probe(ctx context.Context, raw string) error {
 v := oidc.NewVerifier("https://accounts.google.com", oidc.NewRemoteKeySet(ctx,"https://www.googleapis.com/oauth2/v3/certs"), &oidc.Config{ClientID:"probe",SupportedSigningAlgs:[]string{"RS256"}})
 _, err := v.Verify(ctx,raw)
 _ = oauth2.GenerateVerifier()
 return err
}
