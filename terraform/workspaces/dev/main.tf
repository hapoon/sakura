resource "google_artifact_registry_repository" "my-repo" {
  location      = "asia-northeast1"
  repository_id = "sakura"
  description   = "github.com/hapoon/sakura docker repository"
  format        = "DOCKER"
}

