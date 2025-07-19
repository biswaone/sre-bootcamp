# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  # Use Ubuntu 22.04 LTS as base box
  config.vm.box = "ubuntu/jammy64"
  config.vm.network "forwarded_port", guest: 8080, host: 8080
  config.vm.provider "virtualbox" do |vb|
    vb.memory = 4096  # Increase to 4 GB
    vb.cpus = 2       # Use 2 CPUs
  end

  # Install Docker and dependencies
  config.vm.provision "shell", name: "install-dependencies", path: "install-dependencies.sh"

  # Start Terramino (can be rerun with vagrant provision --provision-with start-terramino)
  config.vm.provision "shell", name: "start-sre-bootcamp", inline: <<-SHELL
    cd /home/vagrant/sre-bootcamp
    docker compose up -d
  SHELL
end

