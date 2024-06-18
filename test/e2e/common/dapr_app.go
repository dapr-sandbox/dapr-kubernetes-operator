package common

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/onsi/gomega"
	"github.com/rs/xid"

	"github.com/dapr/kubernetes-operator/test/support"
	"github.com/dapr/kubernetes-operator/test/support/dapr"
	"github.com/lburgazzoli/gomega-matchers/pkg/matchers/jq"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

const (
	IngressPort = 8081
)

func ValidateDaprApp(test support.Test, namespace string) {
	test.T().Helper()

	appName := "testing-app-" + xid.New().String()

	//
	// Install Dapr Test App
	//

	dapr.DeployTestApp(test, appName, namespace)

	test.Eventually(support.Deployment(test, appName, namespace), support.TestTimeoutShort).Should(
		gomega.WithTransform(support.ConditionStatus(appsv1.DeploymentAvailable), gomega.Equal(corev1.ConditionTrue)),
		"Failure checking for App Deployment",
	)
	test.Eventually(support.PodList(test, "app="+appName, namespace), support.TestTimeoutShort).Should(
		gomega.WithTransform(json.Marshal, jq.Match(".items[0].spec.containers | length == 2")),
		"Failure checking for App Pods (sidecar injection)",
	)

	test.Eventually(support.Service(test, appName, namespace), support.TestTimeoutShort).Should(
		gomega.Not(gomega.BeNil()),
		"Failure checking for App Service",
	)
	test.Eventually(support.Ingress(test, appName, namespace), support.TestTimeoutLong).Should(
		gomega.WithTransform(
			support.ExtractFirstLoadBalancerIngressHostname(),
			gomega.Equal("localhost")),
		"Failure to set-up ingress")

	//
	// Test the app
	//

	test.T().Logf("Testing the app with name %s", appName)

	base := fmt.Sprintf("http://localhost:%d/%s", IngressPort, appName)
	value := xid.New().String()

	//nolint:bodyclose
	test.Eventually(dapr.GET(test, base+"/read"), support.TestTimeoutLong).Should(
		gomega.And(
			gomega.HaveHTTPStatus(http.StatusOK),
			gomega.HaveHTTPBody(gomega.Not(jq.Match(`.values | any(. == "%s")`, value))),
		),
		"Failure to read initial values",
	)

	//nolint:bodyclose
	test.Eventually(dapr.POST(test, base+"/write?message="+value, "text/plain", nil), support.TestTimeoutLong).Should(
		gomega.HaveHTTPStatus(http.StatusOK),
		"Failure to store value",
	)

	//nolint:bodyclose
	test.Eventually(dapr.GET(test, base+"/read"), support.TestTimeoutLong).Should(
		gomega.And(
			gomega.HaveHTTPStatus(http.StatusOK),
			gomega.HaveHTTPBody(jq.Match(`.values | any(. == "%s")`, value)),
		),
		"Failure to read final values",
	)
}
