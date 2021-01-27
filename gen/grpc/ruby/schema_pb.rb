# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

require 'google/protobuf'

require 'google/protobuf/struct_pb'
require 'google/protobuf/timestamp_pb'
require 'google/protobuf/any_pb'
require 'google/protobuf/empty_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "meshpaas.Dependency" do
    optional :chart, :string, 1
    optional :version, :string, 2
    optional :repository, :string, 3
  end
  add_message "meshpaas.Maintainer" do
    optional :name, :string, 1
    optional :email, :string, 2
  end
  add_message "meshpaas.Filter" do
    optional :term, :string, 1
    optional :regex, :bool, 2
  end
  add_message "meshpaas.Chart" do
    optional :name, :string, 1
    optional :home, :string, 2
    optional :description, :string, 3
    optional :version, :string, 4
    repeated :sources, :string, 5
    repeated :keywords, :string, 6
    optional :icon, :string, 7
    optional :deprecated, :bool, 8
    repeated :dependencies, :message, 9, "meshpaas.Dependency"
    repeated :maintainers, :message, 10, "meshpaas.Maintainer"
    map :metadata, :string, :string, 11
  end
  add_message "meshpaas.Charts" do
    repeated :charts, :message, 1, "meshpaas.Chart"
  end
  add_message "meshpaas.App" do
    optional :name, :string, 1
    optional :namespace, :string, 2
    optional :release, :message, 5, "meshpaas.Release"
    optional :chart, :message, 20, "meshpaas.Chart"
  end
  add_message "meshpaas.Apps" do
    repeated :apps, :message, 1, "meshpaas.App"
  end
  add_message "meshpaas.Release" do
    optional :version, :uint32, 1
    optional :config, :message, 2, "google.protobuf.Struct"
    optional :notes, :string, 3
    optional :description, :string, 4
    optional :status, :string, 5
    optional :timestamps, :message, 6, "meshpaas.Timestamps"
  end
  add_message "meshpaas.Timestamps" do
    optional :created, :message, 1, "google.protobuf.Timestamp"
    optional :updated, :message, 2, "google.protobuf.Timestamp"
    optional :deleted, :message, 3, "google.protobuf.Timestamp"
  end
  add_message "meshpaas.AppRef" do
    optional :namespace, :string, 1
    optional :name, :string, 2
  end
  add_message "meshpaas.AppInput" do
    optional :namespace, :string, 1
    optional :chart, :string, 2
    optional :app_name, :string, 3
    map :config, :string, :string, 4
  end
  add_message "meshpaas.NamespaceRef" do
    optional :name, :string, 1
  end
  add_message "meshpaas.NamespaceRefs" do
    repeated :namespaces, :message, 1, "meshpaas.NamespaceRef"
  end
end

module Meshpaas
  Dependency = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Dependency").msgclass
  Maintainer = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Maintainer").msgclass
  Filter = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Filter").msgclass
  Chart = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Chart").msgclass
  Charts = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Charts").msgclass
  App = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.App").msgclass
  Apps = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Apps").msgclass
  Release = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Release").msgclass
  Timestamps = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.Timestamps").msgclass
  AppRef = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.AppRef").msgclass
  AppInput = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.AppInput").msgclass
  NamespaceRef = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.NamespaceRef").msgclass
  NamespaceRefs = Google::Protobuf::DescriptorPool.generated_pool.lookup("meshpaas.NamespaceRefs").msgclass
end
