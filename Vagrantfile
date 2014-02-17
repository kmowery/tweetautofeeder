box  = 'precise64'
url  = 'http://cloud-images.ubuntu.com/vagrant/precise/current/precise-server-cloudimg-amd64-vagrant-disk1.box'
sha256 = '79499b35f603be929fd1374eb235c22192aaad72b11f1d101ae73a3293d79fe4'
ip   = '192.168.0.75'
name = 'tweetautofeeder'
ssh_port = 2202

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.define name, primary: true do |machine|
    # Every Vagrant virtual environment requires a box to build off of.
    machine.vm.box = box
    machine.vm.box_url = url
    machine.vm.box_download_checksum = sha256
    machine.vm.box_download_checksum_type = 'sha256'

    # Share SSH locally by default
    machine.vm.network :forwarded_port, guest: 22, host: ssh_port, id: "ssh", auto_correct: true

    # Share some services maybe
    machine.vm.network :forwarded_port, guest: 8080, host: 8080

    # Share an additional folder to the guest VM. The first argument is
    # the path on the host to the actual folder. The second argument is
    # the path on the guest to mount the folder. And the optional third
    # argument is a set of non-required options.
    #machine.vm.synced_folder "../data", "/vagrant_data"
    machine.vm.synced_folder ".", "/home/vagrant/gopath/src/github.com/kmowery/tweetautofeeder/"
    machine.vm.synced_folder "./www", "/usr/share/tweetautofeeder/www", user: "vagrant", group: "vagrant"
    machine.vm.synced_folder "./css", "/usr/share/tweetautofeeder/css", user: "vagrant", group: "vagrant"
    machine.vm.synced_folder "./js", "/usr/share/tweetautofeeder/js", user: "vagrant", group: "vagrant"
    machine.vm.synced_folder "./templates", "/usr/share/tweetautofeeder/templates", user: "vagrant", group: "vagrant"

    # Turn off /vagrant, if you want
    #machine.vm.synced_folder ".", "/vagrant", disabled: true

    # Provision using a shell script
    machine.vm.provision :shell, path: "provision/setup.sh"

    # Provider-specific configuration so you can fine-tune various
    # backing providers for Vagrant. These expose provider-specific options.
    # Example for VirtualBox:
    #
    # machine.vm.provider :virtualbox do |vb|
    #   # Use VBoxManage to customize the VM. For example to change memory:
    #   vb.customize ["modifyvm", :id, "--memory", "1024"]
    # end
  end
end
