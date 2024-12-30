
```
Flags:
  --[no-]help                    Show context-sensitive help (also
                                 try --help-long and --help-man).
                                 ($EXTERNAL_DNS_HELP)
  --[no-]version                 Show application version.
                                 ($EXTERNAL_DNS_VERSION)
  --server=""                    The Kubernetes API server to connect to
                                 (default: auto-detect) ($EXTERNAL_DNS_SERVER)
  --kubeconfig=""                Retrieve target cluster configuration from
                                 a Kubernetes configuration file (default:
                                 auto-detect) ($EXTERNAL_DNS_KUBECONFIG)
  --request-timeout=30s          Request timeout when calling
                                 Kubernetes APIs. 0s means no timeout
                                 ($EXTERNAL_DNS_REQUEST_TIMEOUT)
  --[no-]resolve-service-load-balancer-hostname
                                 Resolve the hostname of LoadBalancer-type
                                 Service object to IP addresses in order to
                                 create DNS A/AAAA records instead of CNAMEs
                                 ($EXTERNAL_DNS_RESOLVE_SERVICE_LOAD_BALANCER_HOSTNAME)
  --cf-api-endpoint=""           The fully-qualified domain name of the
                                 cloud foundry instance you are targeting
                                 ($EXTERNAL_DNS_CF_API_ENDPOINT)
  --cf-username=""               The username to log into the cloud foundry API
                                 ($EXTERNAL_DNS_CF_USERNAME)
  --cf-password=""               The password to log into the cloud foundry API
                                 ($EXTERNAL_DNS_CF_PASSWORD)
  --gloo-namespace=gloo-system ...
                                 The Gloo Proxy namespace; specify multiple
                                 times for multiple namespaces. (default:
                                 gloo-system) ($EXTERNAL_DNS_GLOO_NAMESPACE)
  --skipper-routegroup-groupversion="zalando.org/v1"
                                 The resource version for skipper routegroup
                                 ($EXTERNAL_DNS_SKIPPER_ROUTEGROUP_GROUPVERSION)
  --source=source ...            The resource types that are queried for
                                 endpoints; specify multiple times for multiple
                                 sources (required, options: service, ingress,
                                 node, pod, fake, connector, gateway-httproute,
                                 gateway-grpcroute, gateway-tlsroute,
                                 gateway-tcproute, gateway-udproute,
                                 istio-gateway, istio-virtualservice,
                                 cloudfoundry, contour-httpproxy,
                                 gloo-proxy, crd, empty, skipper-routegroup,
                                 openshift-route, ambassador-host,
                                 kong-tcpingress, f5-virtualserver,
                                 traefik-proxy) ($EXTERNAL_DNS_SOURCE)
  --openshift-router-name=OPENSHIFT-ROUTER-NAME
                                 if source is openshift-route then you
                                 can pass the ingress controller name.
                                 Based on this name external-dns will select
                                 the respective router from the route status
                                 and map that routerCanonicalHostname to the
                                 route host while creating a CNAME record.
                                 ($EXTERNAL_DNS_OPENSHIFT_ROUTER_NAME)
  --namespace=""                 Limit resources queried for endpoints to a
                                 specific namespace (default: all namespaces)
                                 ($EXTERNAL_DNS_NAMESPACE)
  --annotation-filter=""         Filter resources queried for endpoints by
                                 annotation, using label selector semantics
                                 ($EXTERNAL_DNS_ANNOTATION_FILTER)
  --label-filter=""              Filter resources queried for endpoints by label
                                 selector; currently supported by source types
                                 crd, gateway-httproute, gateway-grpcroute,
                                 gateway-tlsroute, gateway-tcproute,
                                 gateway-udproute, ingress, node,
                                 openshift-route, service and ambassador-host
                                 ($EXTERNAL_DNS_LABEL_FILTER)
  --ingress-class=INGRESS-CLASS ...
                                 Require an Ingress to have this class name
                                 (defaults to any class; specify multiple
                                 times to allow more than one class)
                                 ($EXTERNAL_DNS_INGRESS_CLASS)
  --fqdn-template=""             A templated string that's used to generate DNS
                                 names from sources that don't define a hostname
                                 themselves, or to add a hostname suffix when
                                 paired with the fake source (optional). Accepts
                                 comma separated list for multiple global FQDN.
                                 ($EXTERNAL_DNS_FQDN_TEMPLATE)
  --[no-]combine-fqdn-annotation
                                 Combine FQDN template and
                                 Annotations instead of overwriting
                                 ($EXTERNAL_DNS_COMBINE_FQDN_ANNOTATION)
  --[no-]ignore-hostname-annotation
                                 Ignore hostname annotation when
                                 generating DNS names, valid only when
                                 --fqdn-template is set (default: false)
                                 ($EXTERNAL_DNS_IGNORE_HOSTNAME_ANNOTATION)
  --[no-]ignore-ingress-tls-spec
                                 Ignore the spec.tls section in
                                 Ingress resources (default: false)
                                 ($EXTERNAL_DNS_IGNORE_INGRESS_TLS_SPEC)
  --gateway-namespace=GATEWAY-NAMESPACE
                                 Limit Gateways of Route endpoints to a
                                 specific namespace (default: all namespaces)
                                 ($EXTERNAL_DNS_GATEWAY_NAMESPACE)
  --gateway-label-filter=GATEWAY-LABEL-FILTER
                                 Filter Gateways of Route endpoints via
                                 label selector (default: all gateways)
                                 ($EXTERNAL_DNS_GATEWAY_LABEL_FILTER)
  --compatibility=               Process annotation semantics from legacy
                                 implementations (optional, options:
                                 mate, molecule, kops-dns-controller)
                                 ($EXTERNAL_DNS_COMPATIBILITY)
  --[no-]ignore-ingress-rules-spec
                                 Ignore the spec.rules section in
                                 Ingress resources (default: false)
                                 ($EXTERNAL_DNS_IGNORE_INGRESS_RULES_SPEC)
  --[no-]publish-internal-services
                                 Allow external-dns to publish DNS
                                 records for ClusterIP services (optional)
                                 ($EXTERNAL_DNS_PUBLISH_INTERNAL_SERVICES)
  --[no-]publish-host-ip         Allow external-dns to publish host-ip
                                 for headless services (optional)
                                 ($EXTERNAL_DNS_PUBLISH_HOST_IP)
  --[no-]always-publish-not-ready-addresses
                                 Always publish also not ready addresses
                                 for headless services (optional)
                                 ($EXTERNAL_DNS_ALWAYS_PUBLISH_NOT_READY_ADDRESSES)
  --connector-source-server="localhost:8080"
                                 The server to connect for connector source,
                                 valid only when using connector source
                                 ($EXTERNAL_DNS_CONNECTOR_SOURCE_SERVER)
  --crd-source-apiversion="externaldns.k8s.io/v1alpha1"
                                 API version of the CRD for crd source,
                                 e.g. `externaldns.k8s.io/v1alpha1`,
                                 valid only when using crd source
                                 ($EXTERNAL_DNS_CRD_SOURCE_APIVERSION)
  --crd-source-kind="DNSEndpoint"
                                 Kind of the CRD for the crd source in API group
                                 and version specified by crd-source-apiversion
                                 ($EXTERNAL_DNS_CRD_SOURCE_KIND)
  --service-type-filter=SERVICE-TYPE-FILTER ...
                                 The service types to take care about
                                 (default: all, expected: ClusterIP,
                                 NodePort, LoadBalancer or ExternalName)
                                 ($EXTERNAL_DNS_SERVICE_TYPE_FILTER)
  --managed-record-types=A... ...
                                 Record types to manage; specify multiple
                                 times to include many; (default: A, AAAA,
                                 CNAME) (supported records: A, AAAA, CNAME, NS,
                                 SRV, TXT) ($EXTERNAL_DNS_MANAGED_RECORD_TYPES)
  --exclude-record-types=EXCLUDE-RECORD-TYPES ...
                                 Record types to exclude from management;
                                 specify multiple times to exclude many;
                                 (optional) ($EXTERNAL_DNS_EXCLUDE_RECORD_TYPES)
  --default-targets=DEFAULT-TARGETS ...
                                 Set globally default host/IP that will apply as
                                 a target instead of source addresses. Specify
                                 multiple times for multiple targets (optional)
                                 ($EXTERNAL_DNS_DEFAULT_TARGETS)
  --target-net-filter=TARGET-NET-FILTER ...
                                 Limit possible targets by a net filter; specify
                                 multiple times for multiple possible nets
                                 (optional) ($EXTERNAL_DNS_TARGET_NET_FILTER)
  --exclude-target-net=EXCLUDE-TARGET-NET ...
                                 Exclude target nets (optional)
                                 ($EXTERNAL_DNS_EXCLUDE_TARGET_NET)
  --[no-]traefik-disable-legacy  Disable listeners on Resources under
                                 the traefik.containo.us API Group
                                 ($EXTERNAL_DNS_TRAEFIK_DISABLE_LEGACY)
  --[no-]traefik-disable-new     Disable listeners on Resources
                                 under the traefik.io API Group
                                 ($EXTERNAL_DNS_TRAEFIK_DISABLE_NEW)
  --nat64-networks=NAT64-NETWORKS ...
                                 Adding an A record for each AAAA record in
                                 NAT64-enabled networks; specify multiple
                                 times for multiple possible nets (optional)
                                 ($EXTERNAL_DNS_NAT64_NETWORKS)
  --provider=provider            The DNS provider where the DNS records
                                 will be created (required, options: akamai,
                                 alibabacloud, aws, aws-sd, azure, azure-dns,
                                 azure-private-dns, civo, cloudflare,
                                 coredns, designate, digitalocean, dnsimple,
                                 exoscale, gandi, godaddy, google, ibmcloud,
                                 inmemory, linode, ns1, oci, ovh, pdns,
                                 pihole, plural, rfc2136, scaleway, skydns,
                                 tencentcloud, transip, ultradns, webhook)
                                 ($EXTERNAL_DNS_PROVIDER)
  --provider-cache-time=0s       The time to cache the DNS provider record list
                                 requests. ($EXTERNAL_DNS_PROVIDER_CACHE_TIME)
  --domain-filter= ...           Limit possible target zones by a domain suffix;
                                 specify multiple times for multiple domains
                                 (optional) ($EXTERNAL_DNS_DOMAIN_FILTER)
  --exclude-domains= ...         Exclude subdomains (optional)
                                 ($EXTERNAL_DNS_EXCLUDE_DOMAINS)
  --regex-domain-filter=         Limit possible domains and target zones by
                                 a Regex filter; Overrides domain-filter
                                 (optional) ($EXTERNAL_DNS_REGEX_DOMAIN_FILTER)
  --regex-domain-exclusion=      Regex filter that excludes domains and target
                                 zones matched by regex-domain-filter (optional)
                                 ($EXTERNAL_DNS_REGEX_DOMAIN_EXCLUSION)
  --zone-name-filter= ...        Filter target zones by zone domain (For now,
                                 only AzureDNS provider is using this flag);
                                 specify multiple times for multiple zones
                                 (optional) ($EXTERNAL_DNS_ZONE_NAME_FILTER)
  --zone-id-filter= ...          Filter target zones by hosted zone id; specify
                                 multiple times for multiple zones (optional)
                                 ($EXTERNAL_DNS_ZONE_ID_FILTER)
  --google-project=""            When using the Google provider, current
                                 project is auto-detected, when running
                                 on GCP. Specify other project with this.
                                 Must be specified when running outside GCP.
                                 ($EXTERNAL_DNS_GOOGLE_PROJECT)
  --google-batch-change-size=1000
                                 When using the Google provider, set the maximum
                                 number of changes that will be applied in each
                                 batch. ($EXTERNAL_DNS_GOOGLE_BATCH_CHANGE_SIZE)
  --google-batch-change-interval=1s
                                 When using the Google provider,
                                 set the interval between batch changes.
                                 ($EXTERNAL_DNS_GOOGLE_BATCH_CHANGE_INTERVAL)
  --google-zone-visibility=      When using the Google provider,
                                 filter for zones with this visibility
                                 (optional, options: public, private)
                                 ($EXTERNAL_DNS_GOOGLE_ZONE_VISIBILITY)
  --alibaba-cloud-config-file="/etc/kubernetes/alibaba-cloud.json"
                                 When using the Alibaba Cloud provider,
                                 specify the Alibaba Cloud configuration file
                                 (required when --provider=alibabacloud)
                                 ($EXTERNAL_DNS_ALIBABA_CLOUD_CONFIG_FILE)
  --alibaba-cloud-zone-type=     When using the Alibaba Cloud provider,
                                 filter for zones of this type
                                 (optional, options: public, private)
                                 ($EXTERNAL_DNS_ALIBABA_CLOUD_ZONE_TYPE)
  --aws-zone-type=               When using the AWS provider, filter for zones
                                 of this type (optional, options: public,
                                 private) ($EXTERNAL_DNS_AWS_ZONE_TYPE)
  --aws-zone-tags= ...           When using the AWS provider, filter for zones
                                 with these tags ($EXTERNAL_DNS_AWS_ZONE_TAGS)
  --aws-profile= ...             When using the AWS provider, name of the
                                 profile to use ($EXTERNAL_DNS_AWS_PROFILE)
  --aws-assume-role=""           When using the AWS API, assume this IAM role.
                                 Useful for hosted zones in another AWS
                                 account. Specify the full ARN, e.g.
                                 `arn:aws:iam::123455567:role/external-dns`
                                 (optional) ($EXTERNAL_DNS_AWS_ASSUME_ROLE)
  --aws-assume-role-external-id=""
                                 When using the AWS API and assuming a role
                                 then specify this external ID` (optional)
                                 ($EXTERNAL_DNS_AWS_ASSUME_ROLE_EXTERNAL_ID)
  --aws-batch-change-size=1000   When using the AWS provider, set the maximum
                                 number of changes that will be applied in each
                                 batch. ($EXTERNAL_DNS_AWS_BATCH_CHANGE_SIZE)
  --aws-batch-change-size-bytes=32000
                                 When using the AWS provider, set the maximum
                                 byte size that will be applied in each batch.
                                 ($EXTERNAL_DNS_AWS_BATCH_CHANGE_SIZE_BYTES)
  --aws-batch-change-size-values=1000
                                 When using the AWS provider,
                                 set the maximum total record values
                                 that will be applied in each batch.
                                 ($EXTERNAL_DNS_AWS_BATCH_CHANGE_SIZE_VALUES)
  --aws-batch-change-interval=1s
                                 When using the AWS provider,
                                 set the interval between batch changes.
                                 ($EXTERNAL_DNS_AWS_BATCH_CHANGE_INTERVAL)
  --[no-]aws-evaluate-target-health
                                 When using the AWS provider,
                                 set whether to evaluate the health of a
                                 DNS target (default: enabled, disable
                                 with --no-aws-evaluate-target-health)
                                 ($EXTERNAL_DNS_AWS_EVALUATE_TARGET_HEALTH)
  --aws-api-retries=3            When using the AWS API, set the maximum
                                 number of retries before giving up.
                                 ($EXTERNAL_DNS_AWS_API_RETRIES)
  --[no-]aws-prefer-cname        When using the AWS provider, prefer using
                                 CNAME instead of ALIAS (default: disabled)
                                 ($EXTERNAL_DNS_AWS_PREFER_CNAME)
  --aws-zones-cache-duration=0s  When using the AWS provider, set the
                                 zones list cache TTL (0s to disable).
                                 ($EXTERNAL_DNS_AWS_ZONES_CACHE_DURATION)
  --[no-]aws-zone-match-parent   Expand limit possible target by
                                 sub-domains (default: disabled)
                                 ($EXTERNAL_DNS_AWS_ZONE_MATCH_PARENT)
  --[no-]aws-sd-service-cleanup  When using the AWS CloudMap provider,
                                 delete empty Services without
                                 endpoints (default: disabled)
                                 ($EXTERNAL_DNS_AWS_SD_SERVICE_CLEANUP)
  --aws-sd-create-tag=AWS-SD-CREATE-TAG ...
                                 When using the AWS CloudMap provider,
                                 add tag to created services.
                                 The flag can be used multiple times
                                 ($EXTERNAL_DNS_AWS_SD_CREATE_TAG)
  --azure-config-file="/etc/kubernetes/azure.json"
                                 When using the Azure provider,
                                 specify the Azure configuration file
                                 (required when --provider=azure)
                                 ($EXTERNAL_DNS_AZURE_CONFIG_FILE)
  --azure-resource-group=""      When using the Azure provider, override
                                 the Azure resource group to use (optional)
                                 ($EXTERNAL_DNS_AZURE_RESOURCE_GROUP)
  --azure-subscription-id=""     When using the Azure provider, override
                                 the Azure subscription to use (optional)
                                 ($EXTERNAL_DNS_AZURE_SUBSCRIPTION_ID)
  --azure-user-assigned-identity-client-id=""
                                 When using the Azure provider,
                                 override the client id of user assigned
                                 identity in config file (optional)
                                 ($EXTERNAL_DNS_AZURE_USER_ASSIGNED_IDENTITY_CLIENT_ID)
  --azure-zones-cache-duration=0s
                                 When using the Azure provider, set the
                                 zones list cache TTL (0s to disable).
                                 ($EXTERNAL_DNS_AZURE_ZONES_CACHE_DURATION)
  --tencent-cloud-config-file="/etc/kubernetes/tencent-cloud.json"
                                 When using the Tencent Cloud provider,
                                 specify the Tencent Cloud configuration file
                                 (required when --provider=tencentcloud)
                                 ($EXTERNAL_DNS_TENCENT_CLOUD_CONFIG_FILE)
  --tencent-cloud-zone-type=     When using the Tencent Cloud provider,
                                 filter for zones with visibility
                                 (optional, options: public, private)
                                 ($EXTERNAL_DNS_TENCENT_CLOUD_ZONE_TYPE)
  --[no-]cloudflare-proxied      When using the Cloudflare provider, specify
                                 if the proxy mode must be enabled (default:
                                 disabled) ($EXTERNAL_DNS_CLOUDFLARE_PROXIED)
  --cloudflare-dns-records-per-page=100
                                 When using the Cloudflare provider,
                                 specify how many DNS records listed per
                                 page, max possible 5,000 (default: 100)
                                 ($EXTERNAL_DNS_CLOUDFLARE_DNS_RECORDS_PER_PAGE)
  --cloudflare-region-key=CLOUDFLARE-REGION-KEY
                                 When using the Cloudflare provider,
                                 specify the region (default: earth)
                                 ($EXTERNAL_DNS_CLOUDFLARE_REGION_KEY)
  --coredns-prefix="/skydns/"    When using the CoreDNS provider, specify the
                                 prefix name ($EXTERNAL_DNS_COREDNS_PREFIX)
  --akamai-serviceconsumerdomain=""
                                 When using the Akamai provider, specify the
                                 base URL (required when --provider=akamai
                                 and edgerc-path not specified)
                                 ($EXTERNAL_DNS_AKAMAI_SERVICECONSUMERDOMAIN)
  --akamai-client-token=""       When using the Akamai provider, specify the
                                 client token (required when --provider=akamai
                                 and edgerc-path not specified)
                                 ($EXTERNAL_DNS_AKAMAI_CLIENT_TOKEN)
  --akamai-client-secret=""      When using the Akamai provider, specify the
                                 client secret (required when --provider=akamai
                                 and edgerc-path not specified)
                                 ($EXTERNAL_DNS_AKAMAI_CLIENT_SECRET)
  --akamai-access-token=""       When using the Akamai provider, specify the
                                 access token (required when --provider=akamai
                                 and edgerc-path not specified)
                                 ($EXTERNAL_DNS_AKAMAI_ACCESS_TOKEN)
  --akamai-edgerc-path=""        When using the Akamai provider,
                                 specify the .edgerc file path. Path must
                                 be reachable form invocation environment.
                                 (required when --provider=akamai and *-token,
                                 secret serviceconsumerdomain not specified)
                                 ($EXTERNAL_DNS_AKAMAI_EDGERC_PATH)
  --akamai-edgerc-section=""     When using the Akamai provider,
                                 specify the .edgerc file path (Optional
                                 when edgerc-path is specified)
                                 ($EXTERNAL_DNS_AKAMAI_EDGERC_SECTION)
  --oci-config-file="/etc/kubernetes/oci.yaml"
                                 When using the OCI provider, specify the
                                 OCI configuration file (required when
                                 --provider=oci ($EXTERNAL_DNS_OCI_CONFIG_FILE)
  --oci-compartment-ocid=OCI-COMPARTMENT-OCID
                                 When using the OCI provider, specify the
                                 OCID of the OCI compartment containing all
                                 managed zones and records. Required when using
                                 OCI IAM instance principal authentication.
                                 ($EXTERNAL_DNS_OCI_COMPARTMENT_OCID)
  --oci-zone-scope=GLOBAL        When using OCI provider, filter for zones
                                 with this scope (optional, options:
                                 GLOBAL, PRIVATE). Defaults to GLOBAL,
                                 setting to empty value will target both.
                                 ($EXTERNAL_DNS_OCI_ZONE_SCOPE)
  --[no-]oci-auth-instance-principal
                                 When using the OCI provider, specify
                                 whether OCI IAM instance principal
                                 authentication should be used (instead of
                                 key-based auth via the OCI config file).
                                 ($EXTERNAL_DNS_OCI_AUTH_INSTANCE_PRINCIPAL)
  --oci-zones-cache-duration=0s  When using the OCI provider, set the
                                 zones list cache TTL (0s to disable).
                                 ($EXTERNAL_DNS_OCI_ZONES_CACHE_DURATION)
  --inmemory-zone= ...           Provide a list of pre-configured zones for
                                 the inmemory provider; specify multiple
                                 times for multiple zones (optional)
                                 ($EXTERNAL_DNS_INMEMORY_ZONE)
  --ovh-endpoint="ovh-eu"        When using the OVH provider,
                                 specify the endpoint (default: ovh-eu)
                                 ($EXTERNAL_DNS_OVH_ENDPOINT)
  --ovh-api-rate-limit=20        When using the OVH provider,
                                 specify the API request rate limit,
                                 X operations by seconds (default: 20)
                                 ($EXTERNAL_DNS_OVH_API_RATE_LIMIT)
  --pdns-server="http://localhost:8081"
                                 When using the PowerDNS/PDNS provider, specify
                                 the URL to the pdns server (required when
                                 --provider=pdns) ($EXTERNAL_DNS_PDNS_SERVER)
  --pdns-server-id="localhost"   When using the PowerDNS/PDNS provider, specify
                                 the id of the server to retrieve. Should be
                                 `localhost` except when the server is behind a
                                 proxy (optional when --provider=pdns) (default:
                                 localhost) ($EXTERNAL_DNS_PDNS_SERVER_ID)
  --pdns-api-key=""              When using the PowerDNS/PDNS provider,
                                 specify the API key to use to authorize
                                 requests (required when --provider=pdns)
                                 ($EXTERNAL_DNS_PDNS_API_KEY)
  --[no-]pdns-skip-tls-verify    When using the PowerDNS/PDNS provider,
                                 disable verification of any TLS certificates
                                 (optional when --provider=pdns) (default:
                                 false) ($EXTERNAL_DNS_PDNS_SKIP_TLS_VERIFY)
  --ns1-endpoint=""              When using the NS1 provider, specify
                                 the URL of the API endpoint to target
                                 (default: https://api.nsone.net/v1/)
                                 ($EXTERNAL_DNS_NS1_ENDPOINT)
  --[no-]ns1-ignoressl           When using the NS1 provider, specify whether
                                 to verify the SSL certificate (default:
                                 false) ($EXTERNAL_DNS_NS1_IGNORESSL)
  --ns1-min-ttl=NS1-MIN-TTL      Minimal TTL (in seconds) for records.
                                 This value will be used if the provided TTL
                                 for a service/ingress is lower than this.
                                 ($EXTERNAL_DNS_NS1_MIN_TTL)
  --digitalocean-api-page-size=50
                                 Configure the page size used when
                                 querying the DigitalOcean API.
                                 ($EXTERNAL_DNS_DIGITALOCEAN_API_PAGE_SIZE)
  --ibmcloud-config-file="/etc/kubernetes/ibmcloud.json"
                                 When using the IBM Cloud provider,
                                 specify the IBM Cloud configuration
                                 file (required when --provider=ibmcloud
                                 ($EXTERNAL_DNS_IBMCLOUD_CONFIG_FILE)
  --[no-]ibmcloud-proxied        When using the IBM provider, specify if the
                                 proxy mode must be enabled (default: disabled)
                                 ($EXTERNAL_DNS_IBMCLOUD_PROXIED)
  --godaddy-api-key=""           When using the GoDaddy provider, specify the
                                 API Key (required when --provider=godaddy)
                                 ($EXTERNAL_DNS_GODADDY_API_KEY)
  --godaddy-api-secret=""        When using the GoDaddy provider, specify the
                                 API secret (required when --provider=godaddy)
                                 ($EXTERNAL_DNS_GODADDY_API_SECRET)
  --godaddy-api-ttl=GODADDY-API-TTL
                                 TTL (in seconds) for records. This
                                 value will be used if the provided TTL
                                 for a service/ingress is not provided.
                                 ($EXTERNAL_DNS_GODADDY_API_TTL)
  --[no-]godaddy-api-ote         When using the GoDaddy provider,
                                 use OTE api (optional, default:
                                 false, when --provider=godaddy)
                                 ($EXTERNAL_DNS_GODADDY_API_OTE)
  --tls-ca=""                    When using TLS communication, the path
                                 to the certificate authority to verify
                                 server communications (optionally specify
                                 --tls-client-cert for two-way TLS)
                                 ($EXTERNAL_DNS_TLS_CA)
  --tls-client-cert=""           When using TLS communication,
                                 the path to the certificate to present
                                 as a client (not required for TLS)
                                 ($EXTERNAL_DNS_TLS_CLIENT_CERT)
  --tls-client-cert-key=""       When using TLS communication, the path
                                 to the certificate key to use with the
                                 client certificate (not required for TLS)
                                 ($EXTERNAL_DNS_TLS_CLIENT_CERT_KEY)
  --exoscale-apienv="api"        When using Exoscale provider,
                                 specify the API environment (optional)
                                 ($EXTERNAL_DNS_EXOSCALE_APIENV)
  --exoscale-apizone="ch-gva-2"  When using Exoscale provider,
                                 specify the API Zone (optional)
                                 ($EXTERNAL_DNS_EXOSCALE_APIZONE)
  --exoscale-apikey=""           Provide your API Key for the Exoscale provider
                                 ($EXTERNAL_DNS_EXOSCALE_APIKEY)
  --exoscale-apisecret=""        Provide your API Secret for the Exoscale
                                 provider ($EXTERNAL_DNS_EXOSCALE_APISECRET)
  --rfc2136-host=""              When using the RFC2136 provider,
                                 specify the host of the DNS server
                                 ($EXTERNAL_DNS_RFC2136_HOST)
  --rfc2136-port=0               When using the RFC2136 provider,
                                 specify the port of the DNS server
                                 ($EXTERNAL_DNS_RFC2136_PORT)
  --rfc2136-zone=RFC2136-ZONE ...
                                 When using the RFC2136 provider, specify
                                 zone entries of the DNS server to use
                                 ($EXTERNAL_DNS_RFC2136_ZONE)
  --[no-]rfc2136-create-ptr      When using the RFC2136 provider, enable PTR
                                 management ($EXTERNAL_DNS_RFC2136_CREATE_PTR)
  --[no-]rfc2136-insecure        When using the RFC2136 provider, specify
                                 whether to attach TSIG or not (default:
                                 false, requires --rfc2136-tsig-keyname
                                 and rfc2136-tsig-secret)
                                 ($EXTERNAL_DNS_RFC2136_INSECURE)
  --rfc2136-tsig-keyname=""      When using the RFC2136 provider, specify
                                 the TSIG key to attached to DNS messages
                                 (required when --rfc2136-insecure=false)
                                 ($EXTERNAL_DNS_RFC2136_TSIG_KEYNAME)
  --rfc2136-tsig-secret=""       When using the RFC2136 provider, specify the
                                 TSIG (base64) value to attached to DNS messages
                                 (required when --rfc2136-insecure=false)
                                 ($EXTERNAL_DNS_RFC2136_TSIG_SECRET)
  --rfc2136-tsig-secret-alg=""   When using the RFC2136 provider, specify the
                                 TSIG (base64) value to attached to DNS messages
                                 (required when --rfc2136-insecure=false)
                                 ($EXTERNAL_DNS_RFC2136_TSIG_SECRET_ALG)
  --[no-]rfc2136-tsig-axfr       When using the RFC2136 provider, specify the
                                 TSIG (base64) value to attached to DNS messages
                                 (required when --rfc2136-insecure=false)
                                 ($EXTERNAL_DNS_RFC2136_TSIG_AXFR)
  --rfc2136-min-ttl=0s           When using the RFC2136 provider, specify
                                 minimal TTL (in duration format) for records.
                                 This value will be used if the provided TTL
                                 for a service/ingress is lower than this
                                 ($EXTERNAL_DNS_RFC2136_MIN_TTL)
  --[no-]rfc2136-gss-tsig        When using the RFC2136 provider,
                                 specify whether to use secure updates
                                 with GSS-TSIG using Kerberos (default:
                                 false, requires --rfc2136-kerberos-realm,
                                 --rfc2136-kerberos-username,
                                 and rfc2136-kerberos-password)
                                 ($EXTERNAL_DNS_RFC2136_GSS_TSIG)
  --rfc2136-kerberos-username=""
                                 When using the RFC2136 provider with
                                 GSS-TSIG, specify the username of the user
                                 with permissions to update DNS records
                                 (required when --rfc2136-gss-tsig=true)
                                 ($EXTERNAL_DNS_RFC2136_KERBEROS_USERNAME)
  --rfc2136-kerberos-password=""
                                 When using the RFC2136 provider with
                                 GSS-TSIG, specify the password of the user
                                 with permissions to update DNS records
                                 (required when --rfc2136-gss-tsig=true)
                                 ($EXTERNAL_DNS_RFC2136_KERBEROS_PASSWORD)
  --rfc2136-kerberos-realm=""    When using the RFC2136 provider with
                                 GSS-TSIG, specify the realm of the user
                                 with permissions to update DNS records
                                 (required when --rfc2136-gss-tsig=true)
                                 ($EXTERNAL_DNS_RFC2136_KERBEROS_REALM)
  --rfc2136-batch-change-size=50
                                 When using the RFC2136 provider,
                                 set the maximum number of changes
                                 that will be applied in each batch.
                                 ($EXTERNAL_DNS_RFC2136_BATCH_CHANGE_SIZE)
  --[no-]rfc2136-use-tls         When using the RFC2136 provider,
                                 communicate with name server over tls
                                 ($EXTERNAL_DNS_RFC2136_USE_TLS)
  --[no-]rfc2136-skip-tls-verify
                                 When using TLS with the RFC2136 provider,
                                 disable verification of any TLS certificates
                                 ($EXTERNAL_DNS_RFC2136_SKIP_TLS_VERIFY)
  --transip-account=""           When using the TransIP provider, specify the
                                 account name (required when --provider=transip)
                                 ($EXTERNAL_DNS_TRANSIP_ACCOUNT)
  --transip-keyfile=""           When using the TransIP provider,
                                 specify the path to the private key
                                 file (required when --provider=transip)
                                 ($EXTERNAL_DNS_TRANSIP_KEYFILE)
  --pihole-server=""             When using the Pihole provider,
                                 the base URL of the Pihole web server
                                 (required when --provider=pihole)
                                 ($EXTERNAL_DNS_PIHOLE_SERVER)
  --pihole-password=""           When using the Pihole provider, the
                                 password to the server if it is protected
                                 ($EXTERNAL_DNS_PIHOLE_PASSWORD)
  --[no-]pihole-tls-skip-verify  When using the Pihole provider, disable
                                 verification of any TLS certificates
                                 ($EXTERNAL_DNS_PIHOLE_TLS_SKIP_VERIFY)
  --plural-cluster=""            When using the plural provider, specify
                                 the cluster name you're running with
                                 ($EXTERNAL_DNS_PLURAL_CLUSTER)
  --plural-provider=""           When using the plural provider, specify
                                 the provider name you're running with
                                 ($EXTERNAL_DNS_PLURAL_PROVIDER)
  --policy=sync                  Modify how DNS records are synchronized
                                 between sources and providers (default: sync,
                                 options: sync, upsert-only, create-only)
                                 ($EXTERNAL_DNS_POLICY)
  --registry=txt                 The registry implementation to use to keep
                                 track of DNS record ownership (default:
                                 txt, options: txt, noop, dynamodb, aws-sd)
                                 ($EXTERNAL_DNS_REGISTRY)
  --txt-owner-id="default"       When using the TXT or DynamoDB registry, a name
                                 that identifies this instance of ExternalDNS
                                 (default: default) ($EXTERNAL_DNS_TXT_OWNER_ID)
  --txt-prefix=""                When using the TXT registry, a custom string
                                 that's prefixed to each ownership DNS record
                                 (optional). Could contain record type template
                                 like '%{record_type}-prefix-'. Mutual exclusive
                                 with txt-suffix! ($EXTERNAL_DNS_TXT_PREFIX)
  --txt-suffix=""                When using the TXT registry, a custom
                                 string that's suffixed to the host portion
                                 of each ownership DNS record (optional).
                                 Could contain record type template like
                                 '-%{record_type}-suffix'. Mutual exclusive with
                                 txt-prefix! ($EXTERNAL_DNS_TXT_SUFFIX)
  --txt-wildcard-replacement=""  When using the TXT registry, a custom
                                 string that's used instead of an
                                 asterisk for TXT records corresponding
                                 to wildcard DNS records (optional)
                                 ($EXTERNAL_DNS_TXT_WILDCARD_REPLACEMENT)
  --[no-]txt-encrypt-enabled     When using the TXT registry, set if TXT records
                                 should be encrypted before stored (default:
                                 disabled) ($EXTERNAL_DNS_TXT_ENCRYPT_ENABLED)
  --txt-encrypt-aes-key=""       When using the TXT registry, set TXT
                                 record decryption and encryption 32 byte
                                 aes key (required when --txt-encrypt=true)
                                 ($EXTERNAL_DNS_TXT_ENCRYPT_AES_KEY)
  --dynamodb-region=""           When using the DynamoDB registry, the AWS
                                 region of the DynamoDB table (optional)
                                 ($EXTERNAL_DNS_DYNAMODB_REGION)
  --dynamodb-table="external-dns"
                                 When using the DynamoDB registry, the name of
                                 the DynamoDB table (default: "external-dns")
                                 ($EXTERNAL_DNS_DYNAMODB_TABLE)
  --txt-cache-interval=0s        The interval between cache synchronizations
                                 in duration format (default: disabled)
                                 ($EXTERNAL_DNS_TXT_CACHE_INTERVAL)
  --interval=1m0s                The interval between two consecutive
                                 synchronizations in duration format (default:
                                 1m) ($EXTERNAL_DNS_INTERVAL)
  --min-event-sync-interval=5s   The minimum interval between two consecutive
                                 synchronizations triggered from kubernetes
                                 events in duration format (default: 5s)
                                 ($EXTERNAL_DNS_MIN_EVENT_SYNC_INTERVAL)
  --[no-]once                    When enabled, exits the synchronization loop
                                 after the first iteration (default: disabled)
                                 ($EXTERNAL_DNS_ONCE)
  --[no-]dry-run                 When enabled, prints DNS record changes
                                 rather than actually performing them (default:
                                 disabled) ($EXTERNAL_DNS_DRY_RUN)
  --[no-]events                  When enabled, in addition to running every
                                 interval, the reconciliation loop will get
                                 triggered when supported sources change
                                 (default: disabled) ($EXTERNAL_DNS_EVENTS)
  --log-format=text              The format in which log messages are
                                 printed (default: text, options: text, json)
                                 ($EXTERNAL_DNS_LOG_FORMAT)
  --metrics-address=":7979"      Specify where to serve the metrics and
                                 health check endpoint (default: :7979)
                                 ($EXTERNAL_DNS_METRICS_ADDRESS)
  --log-level=info               Set the level of logging. (default: info,
                                 options: panic, debug, info, warning, error,
                                 fatal) ($EXTERNAL_DNS_LOG_LEVEL)
  --webhook-provider-url="http://localhost:8888"
                                 The URL of the remote endpoint
                                 to call for the webhook provider
                                 (default: http://localhost:8888)
                                 ($EXTERNAL_DNS_WEBHOOK_PROVIDER_URL)
  --webhook-provider-read-timeout=5s
                                 The read timeout for the webhook
                                 provider in duration format (default: 5s)
                                 ($EXTERNAL_DNS_WEBHOOK_PROVIDER_READ_TIMEOUT)
  --webhook-provider-write-timeout=10s
                                 The write timeout for the webhook
                                 provider in duration format (default: 10s)
                                 ($EXTERNAL_DNS_WEBHOOK_PROVIDER_WRITE_TIMEOUT)
  --[no-]webhook-server          When enabled, runs as a webhook server
                                 instead of a controller. (default: false).
                                 ($EXTERNAL_DNS_WEBHOOK_SERVER)
```