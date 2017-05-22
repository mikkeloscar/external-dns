/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

type mockCloudFlareClient struct{}

func (m *mockCloudFlareClient) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareClient) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{{ID: "1234567890", Name: "foobar.ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareClient) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudFlareClient) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudFlareClient) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"}, nil
}

func (m *mockCloudFlareClient) ZoneIDByName(zoneName string) (string, error) {
	return "1234567890", nil
}

func (m *mockCloudFlareClient) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

type mockCloudFlareUserDetailsFail struct{}

func (m *mockCloudFlareUserDetailsFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareUserDetailsFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{}, nil
}

func (m *mockCloudFlareUserDetailsFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudFlareUserDetailsFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudFlareUserDetailsFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{}, fmt.Errorf("could not get ID from zone name")
}

func (m *mockCloudFlareUserDetailsFail) ZoneIDByName(zoneName string) (string, error) {
	return "", nil
}

func (m *mockCloudFlareUserDetailsFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

type mockCloudFlareCreateZoneFail struct{}

func (m *mockCloudFlareCreateZoneFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareCreateZoneFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{}, nil
}

func (m *mockCloudFlareCreateZoneFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudFlareCreateZoneFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudFlareCreateZoneFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"}, nil
}

func (m *mockCloudFlareCreateZoneFail) ZoneIDByName(zoneName string) (string, error) {
	return "", nil
}

func (m *mockCloudFlareCreateZoneFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

type mockCloudFlareDNSRecordsFail struct{}

func (m *mockCloudFlareDNSRecordsFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareDNSRecordsFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{}, fmt.Errorf("can not get records from zone")
}
func (m *mockCloudFlareDNSRecordsFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudFlareDNSRecordsFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudFlareDNSRecordsFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"}, nil
}

func (m *mockCloudFlareDNSRecordsFail) ZoneIDByName(zoneName string) (string, error) {
	return "", nil
}

func (m *mockCloudFlareDNSRecordsFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

type mockCloudFlareZoneIDByNameFail struct{}

func (m *mockCloudFlareZoneIDByNameFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareZoneIDByNameFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{}, nil
}

func (m *mockCloudFlareZoneIDByNameFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudFlareZoneIDByNameFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudFlareZoneIDByNameFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{}, nil
}

func (m *mockCloudFlareZoneIDByNameFail) ZoneIDByName(zoneName string) (string, error) {
	return "", fmt.Errorf("no ID for zone found")
}

func (m *mockCloudFlareZoneIDByNameFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

type mockCloudFlareDeleteZoneFail struct{}

func (m *mockCloudFlareDeleteZoneFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareDeleteZoneFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{}, nil
}

func (m *mockCloudFlareDeleteZoneFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudFlareDeleteZoneFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudFlareDeleteZoneFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{}, nil
}

func (m *mockCloudFlareDeleteZoneFail) ZoneIDByName(zoneName string) (string, error) {
	return "1234567890", nil
}

func (m *mockCloudFlareDeleteZoneFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

type mockCloudFlareListZonesFail struct{}

func (m *mockCloudFlareListZonesFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareListZonesFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{}, nil
}

func (m *mockCloudFlareListZonesFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudFlareListZonesFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudFlareListZonesFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{}, nil
}

func (m *mockCloudFlareListZonesFail) ZoneIDByName(zoneName string) (string, error) {
	return "1234567890", nil
}

func (m *mockCloudFlareListZonesFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{}}, fmt.Errorf("no zones available")
}

type mockCloudFlareCreateRecordsFail struct{}

func (m *mockCloudFlareCreateRecordsFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, fmt.Errorf("could not create record")
}

func (m *mockCloudFlareCreateRecordsFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{{ID: "1234567890", Name: "foobar.ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareCreateRecordsFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudFlareCreateRecordsFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudFlareCreateRecordsFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"}, nil
}

func (m *mockCloudFlareCreateRecordsFail) ZoneIDByName(zoneName string) (string, error) {
	return "1234567890", nil
}

func (m *mockCloudFlareCreateRecordsFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{}}, fmt.Errorf("no zones available")
}

type mockCloudFlareDeleteRecordsFail struct{}

func (m *mockCloudFlareDeleteRecordsFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareDeleteRecordsFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{{ID: "1234567890", Name: "foobar.ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareDeleteRecordsFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return nil
}

func (m *mockCloudFlareDeleteRecordsFail) DeleteDNSRecord(zoneID, recordID string) error {
	return fmt.Errorf("could not delete record")
}

func (m *mockCloudFlareDeleteRecordsFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"}, nil
}

func (m *mockCloudFlareDeleteRecordsFail) ZoneIDByName(zoneName string) (string, error) {
	return "1234567890", nil
}

func (m *mockCloudFlareDeleteRecordsFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

type mockCloudFlareUpdateRecordsFail struct{}

func (m *mockCloudFlareUpdateRecordsFail) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return nil, nil
}

func (m *mockCloudFlareUpdateRecordsFail) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return []cloudflare.DNSRecord{{ID: "1234567890", Name: "foobar.ext-dns-test.zalando.to."}}, nil
}

func (m *mockCloudFlareUpdateRecordsFail) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return fmt.Errorf("could not update record")
}

func (m *mockCloudFlareUpdateRecordsFail) DeleteDNSRecord(zoneID, recordID string) error {
	return nil
}

func (m *mockCloudFlareUpdateRecordsFail) UserDetails() (cloudflare.User, error) {
	return cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"}, nil
}

func (m *mockCloudFlareUpdateRecordsFail) ZoneIDByName(zoneName string) (string, error) {
	return "1234567890", nil
}

func (m *mockCloudFlareUpdateRecordsFail) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return []cloudflare.Zone{{Name: "ext-dns-test.zalando.to."}}, nil
}

func TestCloudFlareGetRecordID(t *testing.T) {
	provider := &cloudFlareProvider{
		Client: &mockCloudFlareDNSRecordsFail{},
	}
	zoneID := "12345656790"
	record := cloudflare.DNSRecord{Name: "foobar.ext-dns-test.zalando.to"}
	_, err := provider.getRecordID(zoneID, record)
	if err == nil {
		t.Errorf("exptected error")
	}
	provider.Client = &mockCloudFlareClient{}
	_, err = provider.getRecordID(zoneID, record)
	if err == nil {
		t.Errorf("exptected error")
	}
	record = cloudflare.DNSRecord{Name: "foobar.ext-dns-test.zalando.to."}
	_, err = provider.getRecordID(zoneID, record)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestNewCloudFlareChanges(t *testing.T) {
	action := cloudFlareCreate
	endpoints := []*endpoint.Endpoint{{DNSName: "new", Target: "target"}}
	_ = newCloudFlareChanges(action, endpoints)
}

func TestRecords(t *testing.T) {
	provider := &cloudFlareProvider{
		Client: &mockCloudFlareClient{},
	}
	_, err := provider.Records()
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	provider.Client = &mockCloudFlareDNSRecordsFail{}
	_, err = provider.Records()
	if err == nil {
		t.Errorf("expected to fail")
	}
	provider.Client = &mockCloudFlareListZonesFail{}
	_, err = provider.Records()
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestNewCloudFlareProvider(t *testing.T) {
	_ = os.Setenv("CF_API_KEY", "xxxxxxxxxxxxxxxxx")
	_ = os.Setenv("CF_API_EMAIL", "test@test.com")
	_, err := NewCloudFlareProvider(true)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	_ = os.Unsetenv("CF_API_KEY")
	_ = os.Unsetenv("CF_API_EMAIL")
	_, err = NewCloudFlareProvider(true)
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestApplyChanges(t *testing.T) {
	changes := &plan.Changes{}
	provider := &cloudFlareProvider{
		Client: &mockCloudFlareClient{},
	}
	changes.Create = []*endpoint.Endpoint{{DNSName: "new.ext-dns-test.zalando.to.", Target: "target"}}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "target"}}
	changes.UpdateOld = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "target-old"}}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.zalando.to.", Target: "target-new"}}
	err := provider.ApplyChanges(changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	// empty changes
	changes.Create = []*endpoint.Endpoint{}
	changes.Delete = []*endpoint.Endpoint{}
	changes.UpdateOld = []*endpoint.Endpoint{}
	changes.UpdateNew = []*endpoint.Endpoint{}

	err = provider.ApplyChanges(changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}
