package http

import (
	"bytes"
	"log"
	"net/http"
	"os"
)

// openapiPath is the dev-relative location of the bundled OpenAPI spec,
// resolved against the server's working directory (backend/). go:embed
// cannot reach outside the module tree, so dev-only disk read it is.
const openapiPath = "../api/openapi.yaml"

// devServerURL replaces the spec's relative "url: /api" server so that
// Swagger UI's "Try it out" targets the backend directly on :8090,
// bypassing the Caddy /api reverse-proxy path (which does not strip the
// prefix and would 404 against the backend's /auth/* routes).
const devServerURL = "http://localhost:8090"

// SwaggerHandler serves a single-page Swagger UI (assets from a CDN)
// that loads the OpenAPI spec from /openapi.yaml on the same origin.
// Same-origin => no CORS. Dev-only documentation tooling.
func SwaggerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(swaggerHTML))
	}
}

// OpenAPIHandler serves the bundled OpenAPI spec with its server URL
// rewritten to point directly at the backend (dev only). The source spec
// file on disk is never modified — only the served copy differs.
func OpenAPIHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raw, err := os.ReadFile(openapiPath)
		if err != nil {
			log.Printf("transport: openapi spec read failed: %v", err)
			WriteProblem(w, http.StatusInternalServerError,
				"https://kencleng.dev/problems/internal",
				"Internal Error", "OpenAPI spec is not available.")
			return
		}
		// Rewrite "url: /api" -> the dev server URL so Swagger UI's
		// server selector targets the backend directly.
		rewritten := bytes.Replace(raw,
			[]byte("url: /api"),
			[]byte("url: "+devServerURL),
			1,
		)
		w.Header().Set("Content-Type", "application/yaml")
		_, _ = w.Write(rewritten)
	}
}

const swaggerHTML = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>Kencleng API — Swagger UI</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.18.2/swagger-ui.css" />
  <style>
    body { margin: 0; }
  </style>
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5.18.2/swagger-ui-bundle.js"></script>
  <script src="https://unpkg.com/swagger-ui-dist@5.18.2/swagger-ui-standalone-preset.js"></script>
  <script>
    window.onload = function () {
      window.ui = SwaggerUIBundle({
        url: "/openapi.yaml",
        dom_id: "#swagger-ui",
        deepLinking: true,
        presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
        layout: "BaseLayout",
        tryItOutEnabled: true,
        validatorUrl: null,
        withCredentials: false
      });
    };
  </script>
</body>
</html>
`
