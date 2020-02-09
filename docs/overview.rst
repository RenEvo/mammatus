Overview
========

This private cloud will configure and setup the following components, with some third party software.

.. todo:: Pretty pictures here


Operating System
****************

Unless explicitly called out, `Ubuntu <https://ubuntu.com/>`_ will be used for all cloud infrastructure.

Infrastructure
**************

The cloud will require a certain level of infrastructure, most of these are required with or without hosting a private cloud.

**Network Gateway**

A network gateway is how our network will access the internet and act as a router for the network. 

``iptables`` will be used for configuring the gateway between two ethernet ports.

**Firewall**

A network firewall is the way to control what is allowed to go in and out of the network, on what protocols, and what ports.

``iptables`` will be used for configuring firewall rules.

**DHCP**

Dynamic Host Configuration Protocol is how our servers on the network will receive IP addresses, as well as negotiate the Preboot eXecution Environment (PXE) for the compute nodes.

`dnsmasq <http://www.thekelleys.org.uk/dnsmasq/doc.html>`_ will be used for DHCP

**DNS**

Domain Name System is essentially the phone book for the network resolving names to addresses (i.e. gateway.mammatus==10.0.0.1).

`dnsmasq <http://www.thekelleys.org.uk/dnsmasq/doc.html>`_ will be used for DNS with secondary DNS going to 1.1.1.1.

**TFTP**

Trivial File Transfer Protocol is used by firmware to download images for PXE boot.

`dnsmasq <http://www.thekelleys.org.uk/dnsmasq/doc.html>`_ will be used for TFTP in simple readonly mode.

**Authentication**

Secure Shell (SSH) `keys <https://www.ssh.com/ssh/keygen>`_ will be used for all authentication, and no passwords will be utilized for authenticating across the systems.

**Certificate Authority**

Because the internal services, as well as hosted services, will need Secure Socket Layer (SSL) and Transport Layer Security (TLS) certificates, we will be running a Public Key Infrastructure (PKI).

`smallstep <https://smallstep.com/certificates/>`_ will be used specifically for the Automated Certificate Management Environment (ACME) support and other wide range of integrations.

**Monitoring**

In order to know if things are working, monitoring will need to be in place. A few tools will be used to collect, analyze, report, and alert on metrics.

`TICK Stack <https://www.influxdata.com/time-series-platform/>`_ will be deployed to monitor all cloud infrastructure and provide an alerting framework.

`Grafana <https://grafana.com/>`_ will be deployed to view dashboards and discovery metrics in the system.

`ntopng <https://www.ntop.org/>`_ will be deployed for analyzing traffic on the Network Gateway.

**Command and Control**

Sometimes things need be done across different nodes in the cloud, so we need a way to execute and query the systems in the cloud.

`Serf <https://www.serf.io/>`_ will be deployed and available across the entire cluster.

Cloud Primitives
****************

Cloud primitives define different parts (or features) of a cloud in order to provide an Infrastructure as a Service (IaaS).

**Compute**

Compute resources are provided in order to install and run applications within the cloud. There are many different technologies and virtualizations to fulfill this space.

`lxd <https://linuxcontainers.org/lxd/introduction/>`_ will be used initially as the primary source of allocating compute resources.

**Networking**

For networking on the compute nodes, ``macvlan`` will be used so that newly started containers will use the infrastructure DHCP server to aquire new addresses. This does have a limitation that the container will not be able to communicate with the host hypervisor, but that is a non-issue for my own use cases, and actually somewhat good.

`This blog post <https://blog.simos.info/how-to-make-your-lxd-container-get-ip-addresses-from-your-lan/>`_ has some good steps on how this works and is setup.

.. warning:: Initially network security groups are not a consideration as they are not required for my use case, and really complicate the networking overall.

**Load Balancer**

Load Balancing as a Service (LBaaS) is a service providing load balancing to multiple applications behind it. This can include TCP and/or HTTP applications. Additionally, this allows for multiple applications hosted on the same port (80/443) to run on the same IP address (or different).

`traefix <https://containo.us/traefik/>`_, `KEMP <https://freeloadbalancer.com/>`_, and `gobetween <http://gobetween.io/>`_ are all being considered for this role.

**Block Storage**

Block storage provides a way to allocate block devices (disks) for other systems. The primary use case is for compute instances to have additional storage beyond the flavor sizes, or to provide persistent storage through rebuilds of the instances.

`Linux-IO (LIO) <http://linux-iscsi.org/wiki/LIO>`_ will be used in conjunction with `targetcli <https://linux.die.net/man/8/targetcli>`_ to provide Internet Small Computer Systems Interface (iSCSI) block devices.

**Object Storage**

Object storage is one of those items that most people associate as a *cloud*, being it allows the arbitrary storage and retrieval of *objects* over http. This might include images, documents, backups, disk images, etc... It is really just an all purpose place to storage and distribute files for easy access.

`MinIO <https://min.io>`_ will be used as the Object Store providing an S3 compatable API with policy based bucket access.

**Configuration**

Compute instances need to be customized, having to run external tools after spinning up a new compute instance shouldn't be neccessary. 

`cloud-init <https://cloud-init.io/>`_ will be supported and available in all instances.

**Metadata**

Metadata service is an http API that helps provide remote configurations and information about cloud compute instances essenially providing a *who am i* for instances. This is generally served over http://169.254.169.254 in public clouds. In addition to small metadata api (such as instance-id), the metadata API is also capable of hosting the *user-data* and *meta-data* used in `cloud-init <https://cloud-init.io/>`_ to bootstrap new instances.

``Custom Software`` will be created to fulfill this role, as this is very cloud specific.

**Notification Service**

A Simple Notification Service (SNS) will be provided so that systems can communicate with other systems in the cloud via Pub/Sub over HTTP.

`NATS <https://nats.io>`_ will be used to back a custom SNS HTTP API.

**Queue Service**

A Simple Queue Service (SQS) will be provided so that systems can communicate with other systems with first in first out (FIFO) messaging logic for more reliable Pub/Sub over HTTP.

`NATS Streaming <https://nats.io>`_ will be used to back a custom SQS HTTP API.

**Functions**

Functions as a service (FaaS) will be provided for easy deployment of single functions to execute in the cloud on demand, scheduled, from SNS, or from SQS messages.

`faasd <https://github.com/openfaas/faasd>`_ with `OpenFaaS <https://www.openfaas.com/>`_ will be used to provide this system.

Cloud Services
**************

Additional services will be provided for cloud instance utilization.

**Service Discovery**

`Consul <https://consul.io>`_ will be available for compute instances to join and integrate with. `Consul Agent <https://consul.io/docs/agent>`_ can be installed in compute instances and it will be automatically configured through cloud-init vendor-data to integrate with the cluster.

**Secrets Management**

`Vault <https://vaultproject.io>`_ will be available for compute instances to access and allocate secrets. `Vault Agent <https://www.vaultproject.io/docs/agent/>`_ can be installed in compute instances and will be automatically configured through cloud-init vendor-data to integrate with the secrets transparently and securely.

**Monitoring**

`Telegraf <https://www.influxdata.com/time-series-platform/telegraf/>`_ can be installed in compute instances and will be automatically configured through cloud-init vendor-data to integrate with the Monitoring platform.

**MORE**

Most likely, there will be more... but this is like... a lot to build/document right now.