# policy.rego
package foo

default allow = false
allow {
  input.user == "john"
}