job "sleep" {
  datacenters = ["dc1"]

  group "sleep-alpine" {
    count = 1

    # ephemeral_disk {
    #   size = 10
    # }

    # Disable deployments to reduce scheduling overhead
    update {
      max_parallel = 0
    }


    task "sleep" {
      driver = "docker"

      config {
        image   = "alpine:3.12"
        command = "/bin/sleep"
        args    = ["360000"]
      }

      resources {
        cpu    = 50
        memory = 30
      }

      logs {
        max_files     = 1
        max_file_size = 1
      }
    }
  }
}
