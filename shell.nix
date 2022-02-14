{pkgs ? import <nixpkgs> {}} :
pkgs.mkShell {
    packages = with pkgs; [
        go_1_17
        gopls
    ];
    shellHook = ''
        export PATH=$PATH:$(go env GOPATH)/bin
        export GOFLAGS=-tags=integration
    '';
}
