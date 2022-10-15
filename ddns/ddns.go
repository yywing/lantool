package ddns

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DDNS struct {
	accessKeyId     string
	accessKeySecret string
}

func NewDDNS(ak, sk string) *DDNS {
	return &DDNS{accessKeyId: ak, accessKeySecret: sk}
}

func (s *DDNS) getClient() (*alidns20150109.Client, error) {
	config := &openapi.Config{
		AccessKeyId:     &s.accessKeyId,
		AccessKeySecret: &s.accessKeySecret,
		Endpoint:        tea.String("alidns.cn-hangzhou.aliyuncs.com"),
	}
	return alidns20150109.NewClient(config)
}

// 获取公网 IP
func (s *DDNS) getIP() (string, error) {
	resp, err := http.Get("http://ip.42.pl/raw")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	publicIP := string(content)
	return publicIP, nil
}

func (s *DDNS) Run(domain string, rr string) error {
	client, err := s.getClient()
	if err != nil {
		return nil
	}

	runtime := &util.RuntimeOptions{}

	for {
		publicIP, err := s.getIP()
		if err != nil {
			return err
		}

		describeDomainRecordsRequest := &alidns20150109.DescribeDomainRecordsRequest{
			DomainName: &domain,
			KeyWord:    &rr,
			SearchMode: tea.String("EXACT"),
		}
		resp, err := client.DescribeDomainRecordsWithOptions(describeDomainRecordsRequest, runtime)
		if err != nil {
			return err
		}

		domainCount := len(resp.Body.DomainRecords.Record)
		switch domainCount {
		case 0:
			// 新增
			addDomainRecordRequest := &alidns20150109.AddDomainRecordRequest{
				DomainName: &domain,
				RR:         &rr,
				Type:       tea.String("A"),
				Value:      &publicIP,
			}
			_, err := client.AddDomainRecordWithOptions(addDomainRecordRequest, runtime)
			if err != nil {
				fmt.Printf("create domain err: %v\n", err)
			} else {
				fmt.Printf("create domain: domain %s  rr %s  value %s\n", domain, rr, publicIP)
			}
		case 1:
			// 修改
			domainRecord := resp.Body.DomainRecords.Record[0]
			if *domainRecord.Value == publicIP && *domainRecord.Type == "A" {
				fmt.Printf("equal domain: domain %s  rr %s  value %s\n", domain, rr, publicIP)
			} else {
				updateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{
					RecordId: domainRecord.RecordId,
					RR:       &rr,
					Type:     tea.String("A"),
					Value:    &publicIP,
				}
				_, err = client.UpdateDomainRecordWithOptions(updateDomainRecordRequest, runtime)
				if err != nil {
					fmt.Printf("update domain err: %v\n", err)
				} else {
					fmt.Printf("update domain: domain %s  rr %s  value %s\n", domain, rr, publicIP)
				}
			}
		default:
			return fmt.Errorf("get domain record count error: %d", domainCount)
		}
		<-time.After(60 * time.Second)
	}
}
