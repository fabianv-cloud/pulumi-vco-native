// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export namespace cloudspace {
}

export namespace disk {
}

export namespace ingress {
    export interface HealthCheckArgs {
        interval?: pulumi.Input<number>;
        path?: pulumi.Input<string>;
        port?: pulumi.Input<number>;
        scheme?: pulumi.Input<string>;
        timeout?: pulumi.Input<number>;
    }

    export interface LetsEncryptArgs {
        email: pulumi.Input<string>;
        enabled: pulumi.Input<boolean>;
    }

    export interface OptionsArgs {
        health_check?: pulumi.Input<inputs.ingress.HealthCheckArgs>;
        sticky_session_cookie?: pulumi.Input<inputs.ingress.StickySessionCookieArgs>;
    }

    export interface ReverseProxyBackendArgs {
        options?: pulumi.Input<inputs.ingress.OptionsArgs>;
        scheme: pulumi.Input<string>;
        serverpool_id: pulumi.Input<string>;
        target_port: pulumi.Input<number>;
    }

    export interface ReverseProxyFrontEndArgs {
        domain: pulumi.Input<string>;
        http_port?: pulumi.Input<number>;
        https_port?: pulumi.Input<number>;
        ip_address?: pulumi.Input<string>;
        letsencrypt: pulumi.Input<inputs.ingress.LetsEncryptArgs>;
        scheme: pulumi.Input<string>;
    }

    export interface StickySessionCookieArgs {
        http_only?: pulumi.Input<boolean>;
        name?: pulumi.Input<string>;
        same_site?: pulumi.Input<string>;
        secure?: pulumi.Input<boolean>;
    }

}
