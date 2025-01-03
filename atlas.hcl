data "external_schema" "gorm" {
  program = ["go", "run", "-mod=mod", "./atlasLoader"]
}
env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/17/dev"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
