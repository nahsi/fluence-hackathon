{
  description = "A flake to bootstrap a VM image on Fluence";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.11";
    pre-commit-hooks.url = "github:cachix/pre-commit-hooks.nix";
    comin = {
      url = "github:nlewo/comin";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    ragenix = {
      url = "github:yaxitech/ragenix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs =
    {
      self,
      nixpkgs,
      comin,
      ragenix,
      ...
    }@inputs:
    with inputs;
    let
      system = "x86_64-linux";
      inherit (import ./vars.nix) user;
      pkgs = (import nixpkgs) { inherit system; };
      inherit (nixpkgs) lib;
      specialArgs = {
        inherit
          inputs
          pkgs
          system
          user
          ;
      };
      # function to generate pre-commit-checks
      genChecks =
        system:
        (pre-commit-hooks.lib.${system}.run {
          src = ./.;
          hooks = {
            nixpkgs-fmt.enable = true; # formatter
            statix.enable = true; # linter
            deadnix.enable = true; # linter
          };
        });
    in
    {
      # checks
      checks.${system}.pre-commit-check = genChecks system;

      formatter.x86_64-linux = nixpkgs.legacyPackages.x86_64-linux.nixfmt-rfc-style;

      # See for further options:
      # https://github.com/NixOS/nixpkgs/blob/master/nixos/modules/virtualisation/kubevirt.nix
      nixosConfigurations = {
        init = lib.nixosSystem {
          inherit specialArgs;
          modules = [
            "${nixpkgs}/nixos/modules/virtualisation/kubevirt.nix"
            ./hosts/kubevirt-init
            ragenix.nixosModules.default
            {
              nix.nixPath = [ "nixpkgs=${nixpkgs}" ];
              nix.registry.nixpkgs.flake = nixpkgs;
            }
          ];
        };

        fluence-github-runner = lib.nixosSystem {
          inherit specialArgs;
          modules = [
            comin.nixosModules.comin
            ragenix.nixosModules.default
            ./hosts/fluence-github-runner
          ];
        };
        presentation = lib.nixosSystem {
          inherit specialArgs;
          modules = [
            comin.nixosModules.comin
            ./hosts/presentation
          ];
        };
      };

      packages.${system} = {
        # nix build .#kubevirt-image
        kubevirt-image = self.nixosConfigurations.init.config.system.build.kubevirtImage;
      };
    };
}
