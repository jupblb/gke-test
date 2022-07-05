{ pkgs ? import <nixpkgs> { }
, pkgsLinux ? import <nixpkgs> { system = "x86_64-linux"; }
}:

let test-app = pkgs.buildGoModule {
  postInstall = "cp $src/index.html $out/index.html";
  pname = "test-app";
  src = ./.;
  vendorSha256 = "sha256-PfJNr7t/27PSnwIwFv0kHV3f+er0fpHwqddS8yS7ofo=";
  version = "0.0.0";
};
in
pkgs.dockerTools.buildImage {
  name = "test-app";
  config = {
    Cmd = [ "${test-app}/bin/test-app" ];
    Env = [
      "PORT=8080"
    ];
    ExposedPorts = {
      "8080/tcp" = { };
    };
    WorkingDir = "${test-app}";
  };
}
