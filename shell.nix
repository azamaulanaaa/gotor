{pkgs ? import <nixpkgs> {}} :
pkgs.mkShell {
    packages = with pkgs; [
        go_1_17
        gopls
        python39Packages.grip
    ];
    shellHook = ''
        export PATH=$PATH:$(go env GOPATH)/bin
        export GOFLAGS=-tags=unit,integration
    '';
}
