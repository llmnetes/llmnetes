Yolo operator notes:


on AI models:

- Open ai rate limits API, maybe buy a plan
- Could use more free plans
- Could use open source models? LLAMA from Facebook (check licensing before)
- Bard, hugging face

on AI interpolaribility:

- A bit hard to get a only the manifest we can apply
- Those models will always send text + file + commands
- Sometimes <fields> need to be replaced
- Sometimes need to inject image, change namespace
- Need to ask a response in a specific format see https://github.com/k8sgpt-ai/k8sgpt/blob/main/pkg/ai/prompts.go#L5-L7
- We can ask “don’t explain”  https://github.com/k8sgpt-ai/k8sgpt/blob/main/pkg/ai/prompts.go#L5-L7 
- Request openapiv2 schema https://github.com/k8sgpt-ai/k8sgpt/blob/main/pkg/analysis/analysis.go#L142-L151
- See https://github.com/k8sgpt-ai/k8sgpt/blob/main/pkg/analyzer/pod.go#L47-L103
- Use a cache? to store queries that worked successfully. note that LLMs will likely give a slightly different
  response everytime

- Controller will likely need open RBAC