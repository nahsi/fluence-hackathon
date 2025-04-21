{ inputs, pkgs, ... }:

{
  # Packages
  environment.systemPackages = with pkgs; [
    nixpkgs-fmt
    bash-completion
    git
    vim
    wget
    curl
    btop
    inputs.ragenix.packages.${pkgs.system}.default
  ];
}
