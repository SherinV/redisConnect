Memory used by a redisgraph pod while running with approximately 6000 resources (SNO cluster) was 67Mi. 
```
kubectl top pods
NAME                                                             CPU(cores)   MEMORY(bytes)   
search-operator-5d5b6858c-4487d                                  0m           36Mi            
search-prod-0319f-search-api-79c7578b8b-lh6wn                    7m           39Mi            
search-prod-0319f-search-api-79c7578b8b-wzpnb                    6m           59Mi            
search-prod-0319f-search-collector-5f9d8d684d-8dp79              24m          139Mi           
search-redisgraph-0                                              2m           67Mi   
```

Memory used by a redisgraph pod while running with approximately 1500 resources (secrets only) was 50Mi. 


Rough calculation from redis-data logs

|Cluster| Nodes	|Edges|	Secrets	|Images	|RoleBinding	|ReplicationController|	Pod	|Memory|Secrets+Images|
|------|-------|------|---------|------|-------------|---------------------|-----|------|----------------|
|SNO|6364|3955|1505|281|371|0|213|67Mi|
|SNO - Percent	of Nodes|||0.24|0.04|0.06|0|0.03
|Hub	|11883	|10118	|2328	|320	|593|	22	|474	|124.8Mi|	|
|Hub - Percent	of Nodes||	|0.2|	0.03	|0.05|	0	|0.04	|	|0.23|
|ocpqaaz2	|78245|	40195	|13448|	16472	|9663	|9598	|2050|	821.5Mi|	|
|ocpqaaz2 - Percent	of Nodes|||		0.17|		0.21	|	0.12|		0.12|		0.03	||0.38	|
|ocpqaaz1	|	214643|		39489|		38612|		36496|		29357|		28493|		6478	|	2253.8Mi|	|
|ocpqaaz1 - Percent	of Nodes	|	|		|		0.18|		0.17	|	0.14	|	0.13|		0.03	|	|0.35|
|	ocp	|	114638|		59296|		20773|		18180|		16891|		16370|		2865|		1203.7Mi|	|
|ocp - Percent	of Nodes	|	|	|		0.18|		0.16|		0.15|		0.14|		0.02|	|	0.34|
|Total	|419409	|149098	|75148	|71452	|56498	|54461	|11828	|4.40Gi|	|


Questions?  

Check the ratio of replication controllers (54461) to pods(11828) in total?  
Assumption: Maybe some of the replication controllers are old and can be removed?  
`oc get rc -A  `

Refer:
https://github.com/GoogleCloudPlatform/continuous-deployment-on-kubernetes/issues/29
https://github.com/kubernetes/kubernetes/issues/24330
