variable "user_uuid" {
  description = "The UUID of the user"
  type = string

  validation {
    condition     = can(regex("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$", var.user_uuid))
    error_message = "User UUID must be in the format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  }
}

