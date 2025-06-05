variable "user_uuid" {
  description = "The UUID of the user"
  type = string

  validation {
    condition     = can(regex("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$", var.user_uuid))
    error_message = "User UUID must be in the format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  }
}

variable "bucket_name" {
  type        = string
  description = "The name of the AWS S3 bucket"

  validation {
    condition     = can(regex("^[a-z0-9.-]{3,63}$", var.bucket_name))
    error_message = "Bucket name must be between 3 and 63 characters and contain only lowercase alphanumeric characters, periods, and hyphens."
  }
}

variable "index_html_filepath" {
  description = "The absolute path to the index.html file to be deployed."
  type        = string

  validation {
    condition     = fileexists(var.index_html_filepath)
    error_message = "The specified index.html file path does not exist or is not accessible. Please provide a valid absolute path."
  }
}
variable "error_html_filepath" {
  description = "The absolute path to the error.html file to be deployed."
  type        = string

  validation {
    condition     = fileexists(var.error_html_filepath)
    error_message = "The specified error.html file path does not exist or is not accessible. Please provide a valid absolute path."
  }
}