# terraform-variable-description-coverage

Check coverage of Terraform variable description

## Example

```bash
$ cat test.tf
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
```

```bash
$ ./tfvdc .
test.tf:1:1: variable `foo` does not have description
test.tf:11:1: variable `baz` does not have description
Coverage: 0.67 (2/3)
```
