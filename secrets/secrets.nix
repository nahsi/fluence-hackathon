let
  nahsi-framework = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAILd/6tTC0ZiExgsuvZnJzF32mjFVJBRwZDcUuKb3d5ia";
  nahsi-system76 = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIISV8Fd1iZ8a3Lc9Cb3Ule2M47JGbG8xKMJTSq1ae7ae";
  users = [
    nahsi-framework
    nahsi-system76
  ];

  fluence-github-runner = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIBwZU1ZgFzSlEyHyfPwSeDc1gEp67ctKjr0i6bG5HLSJ";
  systems = [
    fluence-github-runner
  ];
in
{
  "runner-token.age".publicKeys = systems ++ users;
}
