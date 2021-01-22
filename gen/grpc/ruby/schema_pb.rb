# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

require 'google/protobuf'

require 'google/protobuf/struct_pb'
require 'google/protobuf/timestamp_pb'
require 'google/protobuf/any_pb'
require 'google/protobuf/empty_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "meshpaas.AuthnRule" do
    optional :jwks_uri, :string, 1
    optional :issuer, :string, 2
    repeated :audience, :string, 3
    optional :ouput_payload_header, :string, 4
  end
  add_message "meshpaas.AuthzSubject" do
    repeated :allow_issuers, :string, 6
    repeated :allow_roles, :string, 7
    repeated :allow_audience, :string, 8
  end
  add_message "meshpaas.AuthzDestination" do
    repeated :allow_paths, :string, 2
    repeated :allow_hosts, :string, 3
    repeated :allow_methods, :string, 4
    repeated :allow_ports, :string, 5
  end
  add_message "meshpaas.AuthzRule" do
    optional :subject, :message, 1, "meshpaas.AuthzSubject"
    optional :destination, :message, 2, "meshpaas.AuthzDestination"
  end
  add_message "meshpaas.Authz" do
    repeated :rules, :message, 1, "meshpaas.AuthzRule"
  end
  add_message "meshpaas.Authn" do
    repeated :rules, :message, 1, "meshpaas.AuthnRule"
  end
  add_message "meshpaas.SecretInput" do
    optional :name, :string, 1
    optional :type, :enum, 3, "meshpaas.SecretType"
    optional :immutable, :bool, 4
    map :data, :string, :string, 5
  end
  add_message "meshpaas.Secret" do
    optional :name, :string, 1
    optional :type, :enum, 3, "meshpaas.SecretType"
    optional :immutable, :bool, 4
    map :data, :string, :string, 5
  end
  add_message "meshpaas.ServerTLSSettings" do
    optional :https_redirect, :bool, 1
    optional :mode, :enum, 2, "meshpaas.TLSmode"
    optional :secret_name, :string, 3
    repeated :subject_alt_names, :string, 4
    repeated :verify_certificate_spki, :string, 5
    repeated :verify_certificate_hash, :string, 6
    repeated :cipher_suites, :string, 7
  end
  add_message "meshpaas.GatewayListener" do
    optional :port, :uint32, 1
    optional :name, :string, 2
    optional :protocol, :enum, 3, "meshpaas.TransportProtocol"
    repeated :hosts, :string, 4
    optional :tls_config, :message, 5, "meshpaas.ServerTLSSettings"
  end
  add_message "meshpaas.Gateway" do
    optional :name, :string, 1
    repeated :listeners, :message, 3, "meshpaas.GatewayListener"
  end
  add_message "meshpaas.GatewayInput" do
    optional :name, :string, 1
    repeated :listeners, :message, 3, "meshpaas.GatewayListener"
  end
  add_message "meshpaas.HTTPRoute" do
    optional :name, :string, 1
    optional :port, :uint32, 2
    optional :path_prefix, :string, 3
    optional :rewrite_uri, :string, 5
    repeated :allow_origins, :string, 6
    repeated :allow_methods, :string, 7
    repeated :allow_headers, :string, 8
    repeated :expose_headers, :string, 9
    optional :allow_credentials, :bool, 10
  end
  add_message "meshpaas.Routing" do
    optional :gateway, :string, 1
    repeated :hosts, :string, 2
    repeated :http_routes, :message, 4, "meshpaas.HTTPRoute"
  end
  add_message "meshpaas.Container" do
    optional :name, :string, 1
    optional :image, :string, 2
    repeated :args, :string, 3
    map :env, :string, :string, 4
    map :ports, :string, :uint32, 5
  end
  add_message "meshpaas.App" do
    optional :name, :string, 1
    repeated :containers, :message, 3, "meshpaas.Container"
    optional :replicas, :uint32, 8
    optional :routing, :message, 11, "meshpaas.Routing"
    optional :authentication, :message, 12, "meshpaas.Authn"
    optional :authorization, :message, 13, "meshpaas.Authz"
    optional :image_pull_secret, :string, 14
    optional :status, :message, 20, "meshpaas.AppStatus"
  end
  add_message "meshpaas.Task" do
    optional :name, :string, 1
    optional :image_pull_secret, :string, 3
    repeated :containers, :message, 4, "meshpaas.Container"
    optional :schedule, :string, 7
    optional :completions, :uint32, 8
  end
  add_message "meshpaas.TaskInput" do
    optional :name, :string, 1
    optional :image_pull_secret, :string, 3
    repeated :containers, :message, 4, "meshpaas.Container"
    optional :schedule, :string, 7
    optional :completions, :uint32, 8
  end
  add_message "meshpaas.AppInput" do
    optional :name, :string, 1
    repeated :containers, :message, 3, "meshpaas.Container"
    optional :replicas, :uint32, 7
    optional :routing, :message, 10, "meshpaas.Routing"
    optional :authentication, :message, 12, "meshpaas.Authn"
    optional :authorization, :message, 13, "meshpaas.Authz"
    optional :image_pull_secret, :string, 14
  end
  add_message "meshpaas.Ref" do
    optional :name, :string, 1
  end
  add_message "meshpaas.Replica" do
    optional :phase, :string, 1
    optional :condition, :string, 2
    optional :reason, :string, 3
  end
  add_message "meshpaas.AppStatus" do
    repeated :replicas, :message, 1, "meshpaas.Replica"
  end
  add_message "meshpaas.Log" do
    optional :message, :string, 1
  end
  add_message "meshpaas.Apps" do
    repeated :applications, :message, 1, "meshpaas.App"
  end
  add_message "meshpaas.Tasks" do
    repeated :tasks, :message, 1, "meshpaas.Task"
  end
  add_enum "meshpaas.SecretType" do
    value :OPAQUE, 0
    value :TLS_CERT_KEY, 1
    value :DOCKER_CONFIG, 2
  end
  add_enum "meshpaas.TransportProtocol" do
    value :INVALID_PROTOCOL, 0
    value :HTTP, 1
    value :HTTPS, 2
    value :GRPC, 3
    value :HTTP2, 4
    value :MONGO, 5
    value :TCP, 6
    value :TLS, 7
  end
  add_enum "meshpaas.TLSmode" do
    value :PASSTHROUGH, 0
    value :SIMPLE, 1
    value :MUTUAL, 2
    value :AUTO_PASSTHROUGH, 3
    value :ISTIO_MUTUAL, 4
  end
end

module Meshpaas
  AuthnRule = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.AuthnRule").msgclass
  AuthzSubject = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.AuthzSubject").msgclass
  AuthzDestination = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.AuthzDestination").msgclass
  AuthzRule = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.AuthzRule").msgclass
  Authz = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Authz").msgclass
  Authn = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Authn").msgclass
  SecretInput = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.SecretInput").msgclass
  Secret = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Secret").msgclass
  ServerTLSSettings = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.ServerTLSSettings").msgclass
  GatewayListener = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.GatewayListener").msgclass
  Gateway = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Gateway").msgclass
  GatewayInput = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.GatewayInput").msgclass
  HTTPRoute = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.HTTPRoute").msgclass
  Routing = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Routing").msgclass
  Container = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Container").msgclass
  App = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.App").msgclass
  Task = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Task").msgclass
  TaskInput = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.TaskInput").msgclass
  AppInput = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.AppInput").msgclass
  Ref = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Ref").msgclass
  Replica = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Replica").msgclass
  AppStatus = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.AppStatus").msgclass
  Log = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Log").msgclass
  Apps = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Apps").msgclass
  Tasks = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Tasks").msgclass
  SecretType = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.SecretType").enummodule
  TransportProtocol = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.TransportProtocol").enummodule
  TLSmode = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.TLSmode").enummodule
end
