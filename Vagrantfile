$script = <<EOF
set -e

if [[ $(dpkg -l | grep erlang | tail -n 1) == "" ]]; then
  sudo apt-get install -y python-software-properties
  sudo add-apt-repository -y ppa:scottritchie/erlang-backports
  sudo apt-get update
  sudo apt-get -y install erlang
fi

if [[ $(dpkg -l | grep build-essential | tail -n 1) == "" ]]; then
  sudo apt-get -y install build-essential
fi

if [[ $(dpkg -l | grep git-core | tail -n 1) == "" ]]; then
  sudo apt-get -y install git-core
fi

if [ -d teles ]; then
  cd teles && git pull && cd ..
else
  git clone https://github.com/armon/teles
fi

if [ ! -x teles/rel/teles/bin/teles ]; then
  pushd teles
  make deps
  make rel
  popd
fi

ps elf | grep -i teles | awk '{print "kill -9 "$2}' | sh

sudo sudo ./teles/rel/teles/bin/teles start
echo "Done."
EOF

Vagrant.configure("2") do |config|
  config.vm.box = "precise64"
  config.vm.provision :shell, :inline => $script
  config.vm.network :bridged, :bridge => "en0: Wi-Fi (AirPort)"
end
