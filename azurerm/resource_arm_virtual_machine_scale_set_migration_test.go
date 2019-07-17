package azurerm

import (
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestResourceVirtualMachineScaleSetMigrateState(t *testing.T) {
	cases := map[string]struct {
		StateVersion       int
		ID                 string
		InputAttributes    map[string]string
		ExpectedAttributes map[string]string
	}{
		"v1_2_hash": {
			StateVersion: 1,
			ID:           "some_id",
			InputAttributes: map[string]string{
				"os_profile_linux_config.#":                                        "1",
				"os_profile_linux_config.14312344.disable_password_authentication": "true",
				"os_profile_linux_config.14312344.ssh_keys.#":                      "1",
				"os_profile_linux_config.14312344.ssh_keys.0.key_data":             "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCkQWRTmp4UenQJ36usDUE4BxWoHqzGmTohPeb8/N+hpeNH/Q/eLwjZ+1Yjp1OObfliZbqBXdxKPpTlZSVMYJ+F7HlZc0BFYCylgBO9uHu6fRX42jzNa2S9Ho/SIsl5GRmOYK3YyEpdbOHvknCWpbADT8nklLqiYznPNkTAijFHAYcWgZggjQTQqV6YGMUvfovlx479EqrQFap7xUFFbK6lG2pAXdMjfvvivxn8lsP4krbg8yBXjvRaDjpk5YVqyZmYK5EgzcaN0Itj5s1lrCPASTpP2Q3HMB3GS8XYfmZ5JoYyXlp0xn27/m7R5zA3gFTzT8tc6Z5fgcgRqPFzkLHnE+qJkdJBCkUGgsjZKqWKoVqvFSRWEihgvZizFMLIlKaldWTo2arGXSHunPI8gaPDs/9dBKVdA5Zsm1YlJ1o/8CL4e7R10FEnPsVnrCySa1Tt/mmjlH9SRXjWtfTT20FMyOtPhbZJYZnlHBKaW74J42A2oEwaKy6+hpgqExMznA6zTsxA1gnghp+gQdE+ejonCK9Fp5KEGk4zFthNcQrkSSoEPGIlYVHeia4p+hUcfPmpTfHjGcgIXp1ckj5fEmsnBimmEnpOTz+3aXae1uCNkndoGMkF82j4Ti3aMZ7pq5xbhgCP+VjvEdkJpLkzSUCW4n1EOOKj2jrnBfIQQrRoQQ==\n",
				"os_profile_linux_config.14312344.ssh_keys.0.path":                 "/home/ubuntu/.ssh/authorized_keys",
				"name":              "acc-test",
				"network_profile.#": "1",
				"network_profile.123451312.accelerated_networking":                                                 "true",
				"network_profile.123451312.dns_settings.#":                                                         "1",
				"network_profile.123451312.dns_settings.0.dns_servers.#":                                           "0",
				"network_profile.123451312.ip_configuration.#":                                                     "1",
				"network_profile.123451312.ip_configuration.0.application_gateway_backend_address_pool_ids.#":      "0",
				"network_profile.123451312.ip_configuration.0.application_security_group_ids.#":                    "3",
				"network_profile.123451312.ip_configuration.0.application_security_group_ids.2391800824":           "common-asg",
				"network_profile.123451312.ip_configuration.0.application_security_group_ids.3040429762":           "another-asg",
				"network_profile.123451312.ip_configuration.0.application_security_group_ids.617424955":            "one-more-asg",
				"network_profile.123451312.ip_configuration.0.load_balancer_backend_address_pool_ids.#":            "2",
				"network_profile.123451312.ip_configuration.0.load_balancer_backend_address_pool_ids.147218285":    "backend_lb_id",
				"network_profile.123451312.ip_configuration.0.load_balancer_backend_address_pool_ids.4138848892":   "backend_lb_id2",
				"network_profile.123451312.ip_configuration.0.load_balancer_inbound_nat_rules_ids.#":               "0",
				"network_profile.123451312.ip_configuration.0.name":                                                "acc-test",
				"network_profile.123451312.ip_configuration.0.primary":                                             "true",
				"network_profile.123451312.ip_configuration.0.public_ip_address_configuration.#":                   "1",
				"network_profile.123451312.ip_configuration.0.public_ip_address_configuration.0.domain_name_label": "acc-test",
				"network_profile.123451312.ip_configuration.0.public_ip_address_configuration.0.idle_timeout":      "4",
				"network_profile.123451312.ip_configuration.0.public_ip_address_configuration.0.name":              "acc-test",
				"network_profile.123451312.ip_configuration.0.subnet_id":                                           "managment_subnet",
				"network_profile.123451312.ip_forwarding":                                                          "false",
				"network_profile.123451312.name":                                                                   "acc-network",
				"network_profile.123451312.network_security_group_id":                                              "acc-test-nsg",
				"network_profile.123451312.primary":                                                                "true",
				"os_profile.#":                                                                                     "1",
			},
			ExpectedAttributes: map[string]string{
				"os_profile_linux_config.#":                                        "1",
				"os_profile_linux_config.69840937.disable_password_authentication": "true",
				"os_profile_linux_config.69840937.ssh_keys.#":                      "1",
				"os_profile_linux_config.69840937.ssh_keys.0.key_data":             "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCkQWRTmp4UenQJ36usDUE4BxWoHqzGmTohPeb8/N+hpeNH/Q/eLwjZ+1Yjp1OObfliZbqBXdxKPpTlZSVMYJ+F7HlZc0BFYCylgBO9uHu6fRX42jzNa2S9Ho/SIsl5GRmOYK3YyEpdbOHvknCWpbADT8nklLqiYznPNkTAijFHAYcWgZggjQTQqV6YGMUvfovlx479EqrQFap7xUFFbK6lG2pAXdMjfvvivxn8lsP4krbg8yBXjvRaDjpk5YVqyZmYK5EgzcaN0Itj5s1lrCPASTpP2Q3HMB3GS8XYfmZ5JoYyXlp0xn27/m7R5zA3gFTzT8tc6Z5fgcgRqPFzkLHnE+qJkdJBCkUGgsjZKqWKoVqvFSRWEihgvZizFMLIlKaldWTo2arGXSHunPI8gaPDs/9dBKVdA5Zsm1YlJ1o/8CL4e7R10FEnPsVnrCySa1Tt/mmjlH9SRXjWtfTT20FMyOtPhbZJYZnlHBKaW74J42A2oEwaKy6+hpgqExMznA6zTsxA1gnghp+gQdE+ejonCK9Fp5KEGk4zFthNcQrkSSoEPGIlYVHeia4p+hUcfPmpTfHjGcgIXp1ckj5fEmsnBimmEnpOTz+3aXae1uCNkndoGMkF82j4Ti3aMZ7pq5xbhgCP+VjvEdkJpLkzSUCW4n1EOOKj2jrnBfIQQrRoQQ==\n",
				"os_profile_linux_config.69840937.ssh_keys.0.path":                 "/home/ubuntu/.ssh/authorized_keys",
				"name":              "acc-test",
				"network_profile.#": "1",
				"network_profile.2270531745.accelerated_networking":                                                 "true",
				"network_profile.2270531745.dns_settings.#":                                                         "1",
				"network_profile.2270531745.dns_settings.0.dns_servers.#":                                           "0",
				"network_profile.2270531745.ip_configuration.#":                                                     "1",
				"network_profile.2270531745.ip_configuration.0.application_gateway_backend_address_pool_ids.#":      "0",
				"network_profile.2270531745.ip_configuration.0.application_security_group_ids.#":                    "3",
				"network_profile.2270531745.ip_configuration.0.application_security_group_ids.3894836646":           "common-asg",
				"network_profile.2270531745.ip_configuration.0.application_security_group_ids.835205686":            "another-asg",
				"network_profile.2270531745.ip_configuration.0.application_security_group_ids.2321610563":           "one-more-asg",
				"network_profile.2270531745.ip_configuration.0.load_balancer_backend_address_pool_ids.#":            "2",
				"network_profile.2270531745.ip_configuration.0.load_balancer_backend_address_pool_ids.927210270":    "backend_lb_id",
				"network_profile.2270531745.ip_configuration.0.load_balancer_backend_address_pool_ids.3773679477":   "backend_lb_id2",
				"network_profile.2270531745.ip_configuration.0.load_balancer_inbound_nat_rules_ids.#":               "0",
				"network_profile.2270531745.ip_configuration.0.name":                                                "acc-test",
				"network_profile.2270531745.ip_configuration.0.primary":                                             "true",
				"network_profile.2270531745.ip_configuration.0.public_ip_address_configuration.#":                   "1",
				"network_profile.2270531745.ip_configuration.0.public_ip_address_configuration.0.domain_name_label": "acc-test",
				"network_profile.2270531745.ip_configuration.0.public_ip_address_configuration.0.idle_timeout":      "4",
				"network_profile.2270531745.ip_configuration.0.public_ip_address_configuration.0.name":              "acc-test",
				"network_profile.2270531745.ip_configuration.0.subnet_id":                                           "managment_subnet",
				"network_profile.2270531745.ip_forwarding":                                                          "false",
				"network_profile.2270531745.name":                                                                   "acc-network",
				"network_profile.2270531745.network_security_group_id":                                              "acc-test-nsg",
				"network_profile.2270531745.primary":                                                                "true",
				"os_profile.#":                                                                                      "1",
				"id":                                                                                                "some_id",
			},
		},
	}

	for tn, tc := range cases {
		is := &terraform.InstanceState{
			ID:         tc.ID,
			Attributes: tc.InputAttributes,
		}
		is, err := resourceVirtualMachineScaleSetMigrateState(tc.StateVersion, is, nil)

		if err != nil {
			t.Fatalf("bad: %s, err: %#v", tn, err)
		}
		for k, v := range is.Attributes {
			actual := tc.ExpectedAttributes[k]
			if actual != v {
				t.Fatalf("Bad diff Migrate for %q: %q\n\n expected: %q", k, actual, v)
			}
		}
	}
}
