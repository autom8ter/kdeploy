package gql

import (
	"github.com/autom8ter/meshpaas/gen/gql/go/model"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/helpers"
)

func toAPI(input model.APIInput) *meshpaaspb.APIInput {
	var networking = &meshpaaspb.Routing{}
	var containers []*meshpaaspb.Container
	for _, c := range input.Containers {
		ports := []*meshpaaspb.ContainerPort{}
		for _, v := range c.Ports {
			ports = append(ports, &meshpaaspb.ContainerPort{
				Name:   v.Name,
				Number: uint32(v.Number),
				Expose: v.Expose,
			})
		}
		containers = append(containers, &meshpaaspb.Container{
			Name:  c.Name,
			Image: c.Image,
			Args:  c.Args,
			Env:   helpers.ConvertMapS(c.Env),
			Ports: ports,
		})
	}
	for _, r := range input.Routing.HTTPRoutes {
		var (
			pathPrefix       string
			rewriteUri       string
			allowCredentials bool
		)

		networking.HttpRoutes = append(networking.HttpRoutes, &meshpaaspb.HTTPRoute{
			Name:             r.Name,
			Port:             uint32(r.Port),
			PathPrefix:       pathPrefix,
			RewriteUri:       rewriteUri,
			AllowOrigins:     r.AllowOrigins,
			AllowMethods:     r.AllowMethods,
			AllowHeaders:     r.AllowHeaders,
			ExposeHeaders:    r.ExposeHeaders,
			AllowCredentials: allowCredentials,
		})
	}
	return &meshpaaspb.APIInput{
		Name:       input.Name,
		Containers: containers,
		Replicas:   uint32(input.Replicas),
		Routing:    networking,
	}
}

func toSecretType(t model.SecretType) meshpaaspb.SecretType {
	switch t {
	case model.SecretTypeTLSCertKey:
		return meshpaaspb.SecretType_TLS_CERT_KEY
	case model.SecretTypeDockerConfig:
		return meshpaaspb.SecretType_DOCKER_CONFIG
	default:
		return meshpaaspb.SecretType_OPAQUE
	}
}

func toSecret(secret model.SecretInput) *meshpaaspb.SecretInput {
	return &meshpaaspb.SecretInput{
		Name:      secret.Name,
		Type:      toSecretType(secret.Type),
		Immutable: secret.Immutable,
		Data:      helpers.ConvertMapS(secret.Data),
	}
}

func toGateway(gw model.GatewayInput) *meshpaaspb.GatewayInput {
	return nil
}

func toTask(input model.TaskInput) *meshpaaspb.TaskInput {
	var completions uint32
	var containers []*meshpaaspb.Container
	for _, c := range input.Containers {
		var ports []*meshpaaspb.ContainerPort
		for _, v := range c.Ports {
			ports = append(ports, &meshpaaspb.ContainerPort{
				Name:   v.Name,
				Number: uint32(v.Number),
				Expose: v.Expose,
			})
		}
		containers = append(containers, &meshpaaspb.Container{
			Name:  c.Name,
			Image: c.Image,
			Args:  c.Args,
			Env:   helpers.ConvertMapS(c.Env),
			Ports: ports,
		})
	}
	if input.Completions != nil {
		completions = uint32(*input.Completions)
	}
	return &meshpaaspb.TaskInput{
		Name:        input.Name,
		Containers:  containers,
		Schedule:    input.Schedule,
		Completions: completions,
	}
}

func fromTlsMode(mode meshpaaspb.TLSmode) model.TLSmode {
	switch mode {
	case meshpaaspb.TLSmode_AUTO_PASSTHROUGH:
		return model.TLSmodeAutoPassthrough
	case meshpaaspb.TLSmode_ISTIO_MUTUAL:
		return model.TLSmodeIstioMutual
	case meshpaaspb.TLSmode_MUTUAL:
		return model.TLSmodeMutual
	case meshpaaspb.TLSmode_SIMPLE:
		return model.TLSmodeSimple
	default:
		return model.TLSmodeAutoPassthrough
	}
}

func toTlsMode(mode model.TLSmode) meshpaaspb.TLSmode {
	switch mode {
	case model.TLSmodeAutoPassthrough:
		return meshpaaspb.TLSmode_AUTO_PASSTHROUGH
	case model.TLSmodeIstioMutual:
		return meshpaaspb.TLSmode_ISTIO_MUTUAL
	case model.TLSmodeMutual:
		return meshpaaspb.TLSmode_MUTUAL
	case model.TLSmodeSimple:
		return meshpaaspb.TLSmode_SIMPLE
	default:
		return meshpaaspb.TLSmode_PASSTHROUGH
	}
}

func fromTlsConfig(config *meshpaaspb.ServerTLSSettings) *model.ServerTLSSettings {
	return &model.ServerTLSSettings{
		HTTPSRedirect:         config.HttpsRedirect,
		Mode:                  fromTlsMode(config.Mode),
		Secret:                &config.SecretName,
		SubjectAltNames:       config.SubjectAltNames,
		VerifyCertificateSpki: config.VerifyCertificateSpki,
		VerifyCertificateHash: config.VerifyCertificateHash,
		CipherSuites:          config.CipherSuites,
	}
}

func toTlsConfig(config *model.ServerTLSSettings) *meshpaaspb.ServerTLSSettings {
	m := &meshpaaspb.ServerTLSSettings{
		HttpsRedirect:         config.HTTPSRedirect,
		Mode:                  toTlsMode(config.Mode),
		SubjectAltNames:       config.SubjectAltNames,
		VerifyCertificateSpki: config.VerifyCertificateSpki,
		VerifyCertificateHash: config.VerifyCertificateHash,
		CipherSuites:          config.CipherSuites,
	}
	if config.Secret != nil {
		m.SecretName = *config.Secret
	}
	return m
}

func fromProtocol(protocol meshpaaspb.TransportProtocol) model.TransportProtocol {
	switch protocol {
	case meshpaaspb.TransportProtocol_GRPC:
		return model.TransportProtocolGrpc
	case meshpaaspb.TransportProtocol_HTTP:
		return model.TransportProtocolHTTP
	case meshpaaspb.TransportProtocol_HTTPS:
		return model.TransportProtocolHTTPS
	case meshpaaspb.TransportProtocol_HTTP2:
		return model.TransportProtocolHTTP2
	case meshpaaspb.TransportProtocol_MONGO:
		return model.TransportProtocolMongo
	case meshpaaspb.TransportProtocol_TLS:
		return model.TransportProtocolTLS
	default:
		return model.TransportProtocolTCP
	}
}

func toProtocol(protocol model.TransportProtocol) meshpaaspb.TransportProtocol {
	switch protocol {
	case model.TransportProtocolTCP:
		return meshpaaspb.TransportProtocol_TCP
	case model.TransportProtocolTLS:
		return meshpaaspb.TransportProtocol_TLS
	case model.TransportProtocolMongo:
		return meshpaaspb.TransportProtocol_MONGO
	case model.TransportProtocolHTTP:
		return meshpaaspb.TransportProtocol_HTTP
	case model.TransportProtocolHTTPS:
		return meshpaaspb.TransportProtocol_HTTPS
	case model.TransportProtocolHTTP2:
		return meshpaaspb.TransportProtocol_HTTP2
	case model.TransportProtocolGrpc:
		return meshpaaspb.TransportProtocol_GRPC
	default:
		return meshpaaspb.TransportProtocol_TCP
	}
}

func fromGateway(gw *meshpaaspb.Gateway) *model.Gateway {
	g := &model.Gateway{
		Name: gw.Name,
	}
	for _, l := range gw.GetListeners() {
		g.Listeners = append(g.Listeners, &model.GatewayListener{
			Name:      l.Name,
			Port:      int(l.Port),
			Protocol:  fromProtocol(l.Protocol),
			Hosts:     l.Hosts,
			TLSConfig: fromTlsConfig(l.TlsConfig),
		})
	}
	return g
}

func fromAPI(app *meshpaaspb.API) *model.API {
	var (
		status     = &model.APIStatus{}
		containers []*model.Container
	)
	for _, c := range app.Containers {
		var (
			env   map[string]interface{}
			ports []*model.ContainerPort
		)
		if c.Env != nil {
			env = map[string]interface{}{}
			for k, v := range c.Env {
				env[k] = v
			}
		}
		if c.Ports != nil {
			for _, v := range c.Ports {
				ports = append(ports, &model.ContainerPort{
					Name:   v.Name,
					Number: int(v.Number),
					Expose: v.Expose,
				})
			}
		}
		containers = append(containers, &model.Container{
			Name:  c.Name,
			Image: c.Image,
			Args:  c.Args,
			Env:   env,
			Ports: ports,
		})
	}

	for _, r := range app.Status.Replicas {
		status.Replicas = append(status.Replicas, &model.Replica{
			Phase:     r.Phase,
			Condition: r.Condition,
			Reason:    r.Reason,
		})
	}
	return &model.API{
		Name:       app.Name,
		Containers: containers,
		Replicas:   int(app.Replicas),
		Routing:    fromRouting(app.GetRouting()),
		Status:     status,
	}
}

func fromRouting(networking *meshpaaspb.Routing) *model.Routing {
	var routes []*model.HTTPRoute
	for _, r := range networking.GetHttpRoutes() {
		route := &model.HTTPRoute{
			Name:             r.Name,
			Port:             int(r.Port),
			PathPrefix:       helpers.StringPointer(r.PathPrefix),
			RewriteURI:       helpers.StringPointer(r.RewriteUri),
			AllowOrigins:     r.AllowOrigins,
			AllowMethods:     r.AllowMethods,
			AllowHeaders:     r.AllowHeaders,
			ExposeHeaders:    r.ExposeHeaders,
			AllowCredentials: helpers.BoolPointer(r.AllowCredentials),
		}
		if r.Name != "" {
			route.Name = r.Name
		}
		if r.Port != 0 {
			route.Port = int(r.Port)
		}
		routes = append(routes, route)
	}
	r := &model.Routing{
		Hosts:      networking.GetHosts(),
		HTTPRoutes: routes,
	}
	if networking.Gateway != "" {
		r.Gateway = &networking.Gateway
	}
	return r
}

func fromSecretType(t meshpaaspb.SecretType) model.SecretType {
	switch t {
	case meshpaaspb.SecretType_TLS_CERT_KEY:
		return model.SecretTypeTLSCertKey
	case meshpaaspb.SecretType_DOCKER_CONFIG:
		return model.SecretTypeDockerConfig
	default:
		return model.SecretTypeOpaque
	}
}

func fromSecret(secret *meshpaaspb.Secret) *model.Secret {
	return &model.Secret{
		Type:      fromSecretType(secret.Type),
		Name:      secret.Name,
		Immutable: secret.Immutable,
		Data:      helpers.ConvertMap(secret.Data),
	}
}

func fromTask(app *meshpaaspb.Task) *model.Task {
	var containers []*model.Container
	var completions int
	if app.Completions > 0 {
		completions = int(app.Completions)
	}
	for _, c := range app.Containers {
		var (
			env   map[string]interface{}
			ports []*model.ContainerPort
		)
		if c.Env != nil {
			env = map[string]interface{}{}
			for k, v := range c.Env {
				env[k] = v
			}
		}
		if c.Ports != nil {
			var ports []*model.ContainerPort
			for _, v := range c.Ports {
				ports = append(ports, &model.ContainerPort{
					Name:   v.Name,
					Number: int(v.Number),
					Expose: v.Expose,
				})
			}
		}
		containers = append(containers, &model.Container{
			Name:  c.Name,
			Image: c.Image,
			Args:  c.Args,
			Env:   helpers.ConvertMap(c.Env),
			Ports: ports,
		})
	}

	return &model.Task{
		Name:        app.Name,
		Containers:  containers,
		Schedule:    app.Schedule,
		Completions: &completions,
	}
}
