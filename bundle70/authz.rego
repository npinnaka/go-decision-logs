package authz

default allow = false

allow {
    input.role == "admin"
}