package prototype

import "time"

// Need to think about how to structure this. We will try to train our own model and
// see what kind of input output we could have. For now, we will just use the GPT-3
// model from OpenAI.

type KubeAI interface {
	// nginx replicas 5, deploy
	// maybe: UpgradeApp
	DeployApp(description, kind string)
	// vuln, security, k8s version
	RunAnalysis(kind string)
	// need breakdown per kind
	SimulateChaos(kind string)

	// watch and generate data?
	WatchAndLearn(expectedOutcome string, resources []string, duration time.Time)
	// Watch and block if breaking via validating webhook
}

type KubeAIImpl struct {
	// need client

	// need cache

	// need informers
}
