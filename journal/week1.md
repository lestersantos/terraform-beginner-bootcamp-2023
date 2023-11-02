# Terraform Beginner Bootcamp 2023 - Week 1
## Root Module Structure

Our root module structure is as follows:

```
PROJECT_ROOT
│
├── main.tf            # everything else
├── variables.tf       # stores the structure of input variables
├── terraform.tfvars   # the data of variables we want to load into our terraform project
├── providers.tf       # defined required providers and their configuration
├── outputs.tf         # stores our outputs
└── README.md          # required for root modules
```
  
[Standard Module Structure](https://developer.hashicorp.com/terraform/language/modules/develop/structure)

## Terraform and Input Variables

### Terraform Cloud Variables

In terraform we can set two kind of variables:
- Enviroment Variables - those you would set in your bash terminal eg.AWS credentials
- Terraform Variables - those that you would normally set in your tfvars file

We can set Terraform Cloud variables to be sensitive so they are not shown visibliy in the UI.

### Loading Terraform Input Variables
[Terraform Input Variables](https://developer.hashicorp.com/terraform/language/values/variables)

### var flag
We can use the `-var` flag to set an input variable or override a variable in the tfvars file eg. `terraform -var user_uuid="my-user_id"`

### var-file flag

[Variable Definitioins .tfvars Files](https://developer.hashicorp.com/terraform/language/values/variables)

To set lots of variables, it is more convenient to specify their values in a variable definitions file (with a filename ending in either `.tfvars` or `.tfvars.json`) and then specify that file on the
command line with -var-file:

```bash
terraform apply -var-file="testing.tfvars"
```

### terraform.tfvars

This is the default file to load in terraform variables in blunk

### auto.tfvars

Terraform also automatically loads a number of variable definitions files if they are present.

Any files with names ending in `.auto.tfvars` or `.auto.tfvars.json`,
so it is possible to name your file whatever you wish and have Terraform load it
automatically.

All we have to do is provide the file name with `.auto.tfvars` as an extension. eg. `dev.auto.tfvars`

[Terraform .tfvars files: Variables Management](https://spacelift.io/blog/terraform-tfvars)

Variables in the Terraform Cloud workspace and variables provided through the command line 
always overwrite variables with the same key from files ending in `.auto.tfvars`.

[Terraform Cloud Workspaces: variables](https://developer.hashicorp.com/terraform/cloud-docs/workspaces/variables#loading-variables-from-files)

### Order of terraform variables

[Variable definition Precedence](https://developer.hashicorp.com/terraform/language/values/variables#variable-definition-precedence)

Terraform uses the `last` value it finds, overriding any previous values.

Terraform loads variables in the following order, with later sources taking precedence over earlier ones:

- Env variables.
- `terraform.tfvars` file, if present.
- `terraform.tfvars.json` file, if present.
- Any `*.auto.tfvars` or `*.auto.tfvars.json` files, processed in lexical order of their filenames.
- Any `-var` and `-var-file` options on the command line, in the order they are provided. (This includes variables set by a Terraform Cloud worksapce.)

### Order of terraform cloud variables

[Terraform Cloud Workspace Variable: Precedence](https://developer.hashicorp.com/terraform/cloud-docs/workspaces/variables#precedence)

Terraform Cloud prioritizes and overwrites conflicting variables according to the following precedence:

1. Priority Global Variable Sets
2. Priority Project-Scoped Variable Sets
3. Priority Workspace-Scoped Variable Sets
4. Command Line Argument Variables
5. Local Enviroment Variables Prefixed with `TF_VAR_`
6. Workspace-Specific Variables
7. Workspace-Scoped Variable Sets
8. Project-Scoped Variable Sets
9. Global Variable Sets
10. `*.auto.tfvars` Variable Files
 Variables in the Terraform Cloud Workspace and variables provided through the command line always overwrite variables with the same key from files ending in `.auto.tfvars`.
11. `terraform.tfvars` Variable File
 Variables in th `.auto.tfvars` files take precedence over variables in the `terraform.tfvars` file.

### Terraform Cloud Loading Variables form Files

[Loading Variables from files](https://developer.hashicorp.com/terraform/cloud-docs/workspaces/variables/managing-variables#loading-variables-from-files)

You can set Terraform variable values by providing any number of files ending in `.auto.tfvars`
to worksapces that use Terraform 0.10.0 or later. When you trigger a run, Terraform
automatically loads and uses the variables defined in these files. 
If any variable from the worksapce has the same key as a variable in the file, the `workspace` variable overwrites variable from the file.

- Terraform Cloud loads variables from files for each terraform run,
- Those variables does not automatically persist to the TFC workspace
- Those variables ar not display in the Variables section of the UI

## Dealing with Configuration Drift

## What happens if we lose our state file?

If you lose your statefile, you most likely have to tear down all your cloud infraestructure manually.

You can use terraform `import` but it won't work for all cloud resources. You need to check the terraform providers documentation.
For which resources support import.

### Fix Missing Resources with Terraform Import

`terraform import aws_s3_bucket.bucket bucket-name`

```bash
terraform import aws_s3_bucket.example gnk0dzrm7pj69v5td9qmot5rpgbl86dp
```

[Terraform Import](https://developer.hashicorp.com/terraform/language/import)
[Terraform Import AWS S3 Bucket](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket#import)

### Fix Manual Configuration

If someone goes and delete or modifies cloud resources manually through clickops.

If we run Terraform plan it will attempt to put our infraestructure back into
the expected state fixing Configuration Drift.

