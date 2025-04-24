{
  imports = [
    ../hardware-configuration.nix
    ../../system
  ];

  networking = {
    hostName = "presentation";
  };

  programs = {
    direnv = {
      enable = true;
      enableBashIntegration = true;
      nix-direnv.enable = true;
    };
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
}
