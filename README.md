# Dynatrace-counters

Simple golang application to show the misbehavior of Dynatrace prometheus metrics collection for small counter deltas.

## How to run

``` 
# 0. Set up a dynatrace environment, enable Prometheus metrics scraping and deploy dynatrace-operator (with ActiveGate) to a kubernetes cluster, connected to that env.

# 1. Deploy Kubernetes resources
kubectl apply -f k8s.yaml

# 2. The counter of the resource is increased by 1 every minute - verify that by watching the metrics endpoint:
kubectl port-forward svc/dynatrace-counters 8080:80 -n dynatrace-counters

# in different shell:
curl http://localhost:8080/metrics | grep dynatrace
...
# HELP dynatrace_test_counter
# TYPE dynatrace_test_counter counter
dynatrace_test_counter 2

# 3. Find metric 'dynatrace_test_counter' on dynatrace environment and watch it over some minutes.
#    You will notice that the graph in Dynatrace doesn't raise, while the metric value is increased continuously.

# 4. cleanup:
kubectl delete -f k8s.yaml
```