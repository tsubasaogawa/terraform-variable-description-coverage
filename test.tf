variable "foo" {
  type = string
  // no description
}

variable "bar" {
  type        = string
  description = "bar desc"
}

variable "baz" {
  type        = string
  description = "" // empty description
}
