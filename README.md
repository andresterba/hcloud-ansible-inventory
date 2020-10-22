# hcloud-ansible-inventory

This script generates a inventory in the INI format.
It will look like this:

```ini
[hetzner]
server-1 ansible_host=server1.example.com ansible_user=root
server-2 ansible_host=server2.example.com ansible_user=root
server-3 ansible_host=server3.example.com ansible_user=root
```

## Usage

You have to set the enviroment variabel `HETZNER_CLOUD_API_KEY` with your [Hetzner API token](https://docs.hetzner.cloud/#getting-started).

```sh

```