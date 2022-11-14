resource "netbox_ipam_ip_addresses" "ip_test" {
  address = "192.168.56.0/24"
  description = "IP created by terraform"
  status = "active"

  tag {
    name = "tag1"
    slug = "tag1"
  }

  custom_field {
    name = "cf_boolean"
    type = "boolean"
    value = "true"
  }

  custom_field {
    name = "cf_date"
    type = "date"
    value = "2020-12-25"
  }

  custom_field {
    name = "cf_text"
    type = "text"
    value = "some text"
  }

  custom_field {
    name = "cf_integer"
    type = "integer"
    value = "10"
  }

  custom_field {
    name = "cf_selection"
    type = "select"
    value = "1"
  }

  custom_field {
    name = "cf_url"
    type = "url"
    value = "https://github.com"
  }

  custom_field {
    name = "cf_multi_selection"
    type = "multiselect"
    value = jsonencode([
      "0",
      "1"
    ])
  }

  custom_field {
    name = "cf_json"
    type = "json"
    value = jsonencode({
      stringvalue = "string"
      boolvalue = false
      dictionary = {
        numbervalue = 5
      }
    })
  }

  custom_field {
    name = "cf_object"
    type = "object"
    value = data.netbox_dcim_platform.platform_test.id
  }

  custom_field {
    name = "cf_multi_object"
    type = "multiobject"
    value = jsonencode([
      data.netbox_dcim_platform.platform_test.id,
      data.netbox_dcim_platform.platform_test2.id
    ])
  }
}

resource "netbox_ipam_ip_addresses" "ip6_test" {
  address     = "2001:db8::1234/64"
  status      = "active"
  object_id   = netbox_virtualization_interface.interface_test.id
  object_type = netbox_virtualization_interface.interface_test.type
}

resource "netbox_ipam_ip_addresses" "dynamic_ip_from_prefix" {
  prefix = netbox_ipam_prefix.dynamic_prefix_test.id
  description = "Dynamic IP in dynamic prefix created by terraform"
  status = "active"
}

resource "netbox_ipam_ip_addresses" "dynamic_ip_from_ip_range" {
  ip_range = netbox_ipam_ip_range.range_test.id
  description = "Dynamic IP in IP range created by terraform"
  status = "active"
}
