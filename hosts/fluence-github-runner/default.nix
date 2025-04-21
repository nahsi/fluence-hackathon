{ config, ... }:
{
  imports = [
    ../hardware-configuration.nix
    ../../system
    ../../secrets
  ];

  networking = {
    hostName = "fluence-github-runner";
  };

  # pull configuration from a repo
  services.comin = {
    enable = true;
    remotes = [
      {
        name = "origin";
        url = "https://github.com/nahsi/fluence-hackathon.git";
        branches.main.name = "master";
      }
    ];
  };

  services.github-runners = {
    fluence-hackathon = {
      enable = true;
      name = "fluence-runner";
      tokenFile = "${config.age.secrets.runner-token.path}";
      url = "https://github.com/nahsi/fluence-hackathon";
    };
  };
}
