# policy.rego
package authz

default allow = false
allow {
  input.user == "john"
}