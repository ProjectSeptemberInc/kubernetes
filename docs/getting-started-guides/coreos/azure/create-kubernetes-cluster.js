#!/usr/bin/env node

var azure = require('./lib/azure_wrapper.js');
var kube = require('./lib/deployment_logic/kubernetes.js');
var service_name = process.env['AZ_SERVICE_NAME'] || 'kube';
var etcd_instances = process.env['AZ_ETCD_INSTANCES'] || 3;
var kube_instances = process.env['AZ_KUBE_INSTANCES'] || 3;

azure.create_config(service_name, { 'etcd': etcd_instances, 'kube': kube_instances });

azure.run_task_queue([
  azure.queue_default_network(),
  azure.queue_storage_if_needed(),
  azure.queue_machines('etcd', 'stable',
    kube.create_etcd_cloud_config),
  azure.queue_machines('kube', 'stable',
    kube.create_node_cloud_config),
]);
