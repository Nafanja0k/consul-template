package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestVaultConfig_Copy(t *testing.T) {
	cases := []struct {
		name string
		a    *VaultConfig
	}{
		{
			"nil",
			nil,
		},
		{
			"empty",
			&VaultConfig{},
		},
		{
			"same_enabled",
			&VaultConfig{
				Address:     String("address"),
				Enabled:     Bool(true),
				RenewToken:  Bool(true),
				SSL:         &SSLConfig{Enabled: Bool(true)},
				Token:       String("token"),
				UnwrapToken: Bool(true),
			},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d_%s", i, tc.name), func(t *testing.T) {
			r := tc.a.Copy()
			if !reflect.DeepEqual(tc.a, r) {
				t.Errorf("\nexp: %#v\nact: %#v", tc.a, r)
			}
		})
	}
}

func TestVaultConfig_Merge(t *testing.T) {
	cases := []struct {
		name string
		a    *VaultConfig
		b    *VaultConfig
		r    *VaultConfig
	}{
		{
			"nil_a",
			nil,
			&VaultConfig{},
			&VaultConfig{},
		},
		{
			"nil_b",
			&VaultConfig{},
			nil,
			&VaultConfig{},
		},
		{
			"nil_both",
			nil,
			nil,
			nil,
		},
		{
			"empty",
			&VaultConfig{},
			&VaultConfig{},
			&VaultConfig{},
		},
		{
			"enabled_overrides",
			&VaultConfig{Enabled: Bool(true)},
			&VaultConfig{Enabled: Bool(false)},
			&VaultConfig{Enabled: Bool(false)},
		},
		{
			"enabled_empty_one",
			&VaultConfig{Enabled: Bool(true)},
			&VaultConfig{},
			&VaultConfig{Enabled: Bool(true)},
		},
		{
			"enabled_empty_two",
			&VaultConfig{},
			&VaultConfig{Enabled: Bool(true)},
			&VaultConfig{Enabled: Bool(true)},
		},
		{
			"enabled_same",
			&VaultConfig{Enabled: Bool(true)},
			&VaultConfig{Enabled: Bool(true)},
			&VaultConfig{Enabled: Bool(true)},
		},
		{
			"address_overrides",
			&VaultConfig{Address: String("address")},
			&VaultConfig{Address: String("")},
			&VaultConfig{Address: String("")},
		},
		{
			"address_empty_one",
			&VaultConfig{Address: String("address")},
			&VaultConfig{},
			&VaultConfig{Address: String("address")},
		},
		{
			"address_empty_two",
			&VaultConfig{},
			&VaultConfig{Address: String("address")},
			&VaultConfig{Address: String("address")},
		},
		{
			"address_same",
			&VaultConfig{Address: String("address")},
			&VaultConfig{Address: String("address")},
			&VaultConfig{Address: String("address")},
		},
		{
			"token_overrides",
			&VaultConfig{Token: String("token")},
			&VaultConfig{Token: String("")},
			&VaultConfig{Token: String("")},
		},
		{
			"token_empty_one",
			&VaultConfig{Token: String("token")},
			&VaultConfig{},
			&VaultConfig{Token: String("token")},
		},
		{
			"token_empty_two",
			&VaultConfig{},
			&VaultConfig{Token: String("token")},
			&VaultConfig{Token: String("token")},
		},
		{
			"token_same",
			&VaultConfig{Token: String("token")},
			&VaultConfig{Token: String("token")},
			&VaultConfig{Token: String("token")},
		},
		{
			"unwrap_token_overrides",
			&VaultConfig{UnwrapToken: Bool(true)},
			&VaultConfig{UnwrapToken: Bool(false)},
			&VaultConfig{UnwrapToken: Bool(false)},
		},
		{
			"unwrap_token_empty_one",
			&VaultConfig{UnwrapToken: Bool(true)},
			&VaultConfig{},
			&VaultConfig{UnwrapToken: Bool(true)},
		},
		{
			"unwrap_token_empty_two",
			&VaultConfig{},
			&VaultConfig{UnwrapToken: Bool(true)},
			&VaultConfig{UnwrapToken: Bool(true)},
		},
		{
			"unwrap_token_same",
			&VaultConfig{UnwrapToken: Bool(true)},
			&VaultConfig{UnwrapToken: Bool(true)},
			&VaultConfig{UnwrapToken: Bool(true)},
		},
		{
			"renew_token_overrides",
			&VaultConfig{RenewToken: Bool(true)},
			&VaultConfig{RenewToken: Bool(false)},
			&VaultConfig{RenewToken: Bool(false)},
		},
		{
			"renew_token_empty_one",
			&VaultConfig{RenewToken: Bool(true)},
			&VaultConfig{},
			&VaultConfig{RenewToken: Bool(true)},
		},
		{
			"renew_token_empty_two",
			&VaultConfig{},
			&VaultConfig{RenewToken: Bool(true)},
			&VaultConfig{RenewToken: Bool(true)},
		},
		{
			"renew_token_same",
			&VaultConfig{RenewToken: Bool(true)},
			&VaultConfig{RenewToken: Bool(true)},
			&VaultConfig{RenewToken: Bool(true)},
		},
		{
			"ssl_overrides",
			&VaultConfig{SSL: &SSLConfig{Enabled: Bool(true)}},
			&VaultConfig{SSL: &SSLConfig{Enabled: Bool(false)}},
			&VaultConfig{SSL: &SSLConfig{Enabled: Bool(false)}},
		},
		{
			"ssl_empty_one",
			&VaultConfig{SSL: &SSLConfig{Enabled: Bool(true)}},
			&VaultConfig{},
			&VaultConfig{SSL: &SSLConfig{Enabled: Bool(true)}},
		},
		{
			"ssl_empty_two",
			&VaultConfig{},
			&VaultConfig{SSL: &SSLConfig{Enabled: Bool(true)}},
			&VaultConfig{SSL: &SSLConfig{Enabled: Bool(true)}},
		},
		{
			"ssl_same",
			&VaultConfig{SSL: &SSLConfig{Enabled: Bool(true)}},
			&VaultConfig{SSL: &SSLConfig{Enabled: Bool(true)}},
			&VaultConfig{SSL: &SSLConfig{Enabled: Bool(true)}},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d_%s", i, tc.name), func(t *testing.T) {
			r := tc.a.Merge(tc.b)
			if !reflect.DeepEqual(tc.r, r) {
				t.Errorf("\nexp: %#v\nact: %#v", tc.r, r)
			}
		})
	}
}

func TestVaultConfig_Finalize(t *testing.T) {
	cases := []struct {
		name string
		i    *VaultConfig
		r    *VaultConfig
	}{
		{
			"empty",
			&VaultConfig{},
			&VaultConfig{
				Address:    String(""),
				Enabled:    Bool(false),
				RenewToken: Bool(DefaultVaultRenewToken),
				SSL: &SSLConfig{
					CaCert:     String(""),
					CaPath:     String(""),
					Cert:       String(""),
					Enabled:    Bool(false),
					Key:        String(""),
					ServerName: String(""),
					Verify:     Bool(true),
				},
				Token:       String(""),
				UnwrapToken: Bool(DefaultVaultUnwrapToken),
			},
		},
		{
			"with_address",
			&VaultConfig{
				Address: String("address"),
			},
			&VaultConfig{
				Address:    String("address"),
				Enabled:    Bool(true),
				RenewToken: Bool(DefaultVaultRenewToken),
				SSL: &SSLConfig{
					CaCert:     String(""),
					CaPath:     String(""),
					Cert:       String(""),
					Enabled:    Bool(false),
					Key:        String(""),
					ServerName: String(""),
					Verify:     Bool(true),
				},
				Token:       String(""),
				UnwrapToken: Bool(DefaultVaultUnwrapToken),
			},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d_%s", i, tc.name), func(t *testing.T) {
			tc.i.Finalize()
			if !reflect.DeepEqual(tc.r, tc.i) {
				t.Errorf("\nexp: %#v\nact: %#v", tc.r, tc.i)
			}
		})
	}
}