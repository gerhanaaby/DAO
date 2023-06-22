package utils

func GetEndPoint(ep string, key string, auth string) []string {
	var vals []string
	urls := make(map[string]string)

	urls[`onefcc`] = `https://apidev.banksinarmas.com/internal/kyc/onefcc/v1.0/screening`
	urls[`onefcc-key`] = `99c2db64-2956-49d5-b1f0-414686b35946`
	urls["onefcc-auth"] = ""

	urls[`synccif`] = `https://apidev.banksinarmas.com/internal/bank-services/Customer/v1.0/syncCIF`
	urls[`synccif-key`] = `f869b089-f0d5-4382-a81f-80444342abc9`
	urls["synccif-auth"] = ""

	urls[`csr20`] = `http://10.22.102.18/PegaCSR20/api/S1WS/callWSViaMQ`
	urls[`csr20-key`] = ""
	urls["csr20-auth"] = ""

	urls[`eformcentral`] = `http://10.32.1.77:8080/EFormCentral/rest/formdata/action/applicationDone`
	urls[`eformcentral-key`] = ""
	urls["eformcentral-auth"] = "Basic ZUZvcm1TeXN0ZW06ZUYwcm1DM250cjRs"

	vals = append(vals, urls[ep], urls[key])

	return vals
}
