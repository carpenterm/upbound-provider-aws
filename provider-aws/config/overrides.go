/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/config"
	"github.com/upbound/upjet/pkg/types/comments"
	"github.com/upbound/upjet/pkg/types/name"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// RegionAddition adds region to the spec of all resources except iam group which
// does not have a region notion.
func RegionAddition() config.ResourceOption {
	return func(r *config.Resource) {
		if r.ShortGroup == "iam" {
			return
		}
		c := "Region is the region you'd like your resource to be created in.\n"
		comment, err := comments.New(c, comments.WithTFTag("-"))
		if err != nil {
			panic(errors.Wrap(err, "cannot build comment for region"))
		}
		r.TerraformResource.Schema["region"] = &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: comment.String(),
		}
	}
}

// TagsAllRemoval removes the tags_all field that is used only in tfstate to
// accumulate provider-wide default tags in TF, which is not something we support.
// So, we don't need it as a parameter while "tags" is already in place.
func TagsAllRemoval() config.ResourceOption {
	return func(r *config.Resource) {
		if t, ok := r.TerraformResource.Schema["tags_all"]; ok {
			t.Computed = true
			t.Optional = false
		}
	}
}

// IdentifierAssignedByAWS will work for all AWS types because even if the ID
// is assigned by user, we'll see it in the TF State ID.
// The resource-specific configurations should override this whenever possible.
func IdentifierAssignedByAWS() config.ResourceOption {
	return func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	}
}

// NamePrefixRemoval makes sure we remove name_prefix from all since it is mostly
// for Terraform functionality.
func NamePrefixRemoval() config.ResourceOption {
	return func(r *config.Resource) {
		for _, f := range r.ExternalName.OmittedFields {
			if f == "name_prefix" {
				return
			}
		}
		r.ExternalName.OmittedFields = append(r.ExternalName.OmittedFields, "name_prefix")
	}
}

// KnownReferencers adds referencers for fields that are known and common among
// more than a few resources.
func KnownReferencers() config.ResourceOption { //nolint:gocyclo
	return func(r *config.Resource) {
		for k, s := range r.TerraformResource.Schema {
			// We shouldn't add referencers for status fields and sensitive fields
			// since they already have secret referencer.
			if (s.Computed && !s.Optional) || s.Sensitive {
				continue
			}
			switch {
			case strings.HasSuffix(k, "role_arn"):
				r.References[k] = config.Reference{
					Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role",
					Extractor: common.PathARNExtractor,
				}
			case strings.HasSuffix(k, "security_group_ids"):
				r.References[k] = config.Reference{
					Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.SecurityGroup",
					RefFieldName:      strings.TrimSuffix(name.NewFromSnake(k).Camel, "s") + "Refs",
					SelectorFieldName: strings.TrimSuffix(name.NewFromSnake(k).Camel, "s") + "Selector",
				}
			case r.ShortGroup == "glue" && k == "database_name":
				r.References["database_name"] = config.Reference{
					Type: "github.com/upbound/official-providers/provider-aws/apis/glue/v1beta1.CatalogDatabase",
				}
			}
			switch k {
			case "vpc_id":
				r.References["vpc_id"] = config.Reference{
					Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.VPC",
					RefFieldName:      "VpcIdRef",
					SelectorFieldName: "VpcIdSelector",
				}
			case "subnet_ids":
				r.References["subnet_ids"] = config.Reference{
					Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Subnet",
					RefFieldName:      "SubnetIdRefs",
					SelectorFieldName: "SubnetIdSelector",
				}
			case "subnet_id":
				r.References["subnet_id"] = config.Reference{
					Type: "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Subnet",
				}
			case "iam_roles":
				r.References["iam_roles"] = config.Reference{
					Type:              "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role",
					RefFieldName:      "IAMRoleRefs",
					SelectorFieldName: "IAMRoleSelector",
				}
			case "security_group_id":
				r.References["security_group_id"] = config.Reference{
					Type: "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.SecurityGroup",
				}
			case "kms_key_id":
				r.References["kms_key_id"] = config.Reference{
					Type: "github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key",
				}
			case "kms_key_arn":
				r.References["kms_key_arn"] = config.Reference{
					Type: "github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key",
				}
			case "kms_key":
				r.References["kms_key"] = config.Reference{
					Type: "github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key",
				}
			}
		}
	}
}

// AddExternalTagsField adds ExternalTagsFieldName configuration for resources that have tags field.
func AddExternalTagsField() config.ResourceOption {
	return func(r *config.Resource) {
		if s, ok := r.TerraformResource.Schema["tags"]; ok && s.Type == schema.TypeMap {
			r.InitializerFns = append(r.InitializerFns, config.TagInitializer)
		}
	}
}
