/*
IBM Confidential
OCO Source Materials
(C) Copyright IBM Corporation 2019 All Rights Reserved
The source code for this program is not published or otherwise divested of its trade secrets, irrespective of what has been deposited with the U.S. Copyright Office.
*/
// Copyright (c) 2020 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/golang/glog"
	"github.com/open-cluster-management/redisConnect/pkg/config"
	"github.com/open-cluster-management/redisConnect/pkg/dbconnector"
	"github.com/open-cluster-management/redisConnect/pkg/handlers"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	// parse flags
	flag.Parse()
	// Glog by default logs to a file. Change it so that by default it all goes to stderr. (no option for stdout).
	err := flag.Lookup("logtostderr").Value.Set("true")
	if err != nil {
		fmt.Println("Error setting default flag:", err) // Uses fmt.Println in case something is wrong with glog args
		os.Exit(1)
		glog.Fatal("Error setting default flag: ", err)
	}
	defer glog.Flush() // This should ensure that everything makes it out on to the console if the program crashes.

	glog.Info("Starting redisConnect")
	if commit, ok := os.LookupEnv("VCS_REF"); ok {
		glog.Info("Built from git commit: ", commit)
	}
	ns := config.Cfg.Namespace
	fmt.Println("Namespace present:", ns)
	opts := v1.ListOptions{LabelSelector: "app=search-prod, component=redisgraph"}
	svcList, err := config.GetKubeClient().CoreV1().Services(ns).List(opts)
	if err == nil {
		for _, svc := range svcList.Items {
			if svc.Name != "" && strings.Contains(svc.Name, "search-redisgraph") {
				// search-prod-0319f-search-redisgraph.open-cluster-management.svc
				redisHost := svc.Name + "." + ns + "." + "svc"

				os.Setenv("REDIS_HOST", redisHost)
				fmt.Println("Set REDIS_HOST")

				config.InitConfig()
				dbconnector.InitConfig()
			}
		}
	} else {
		fmt.Println("Error listing services", err)
	}
	dbconnector.GetKinds()
	handlers.ClusterStats()

	// go dbconnector.RedisWatcher()
	// // Watch clusters and sync status to Redis.
	// go clustermgmt.WatchClusters()

	// Run routine to build intercluster edges

	// router := mux.NewRouter()

	// router.HandleFunc("/liveness", handlers.LivenessProbe).Methods("GET")
	// router.HandleFunc("/readiness", handlers.ReadinessProbe).Methods("GET")
	// router.HandleFunc("/aggregator/clusters/{id}/sync", handlers.SyncResources).Methods("POST")

	// Configure TLS
	// cfg := &tls.Config{
	// 	MinVersion:               tls.VersionTLS12,
	// 	CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
	// 	PreferServerCipherSuites: true,
	// 	CipherSuites: []uint16{
	// 		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	// 	},
	// }
	// srv := &http.Server{
	// 	Addr:              config.Cfg.AggregatorAddress,
	// 	Handler:           router,
	// 	TLSConfig:         cfg,
	// 	ReadHeaderTimeout: time.Duration(config.Cfg.HTTPTimeout) * time.Millisecond,
	// 	ReadTimeout:       time.Duration(config.Cfg.HTTPTimeout) * time.Millisecond,
	// 	WriteTimeout:      time.Duration(config.Cfg.HTTPTimeout) * time.Millisecond,
	// 	TLSNextProto:      make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	// }

	// glog.Info("Listening on: ", config.Cfg.AggregatorAddress)
	// log.Fatal(srv.ListenAndServeTLS("./sslcert/tls.crt", "./sslcert/tls.key"),
	// 	" Use ./setup.sh to generate certificates for local development.")
}
