import radius as radius

param magpieimage string = 'radiusdev.azurecr.io/magpiego:latest'

param environment string

param location string = resourceGroup().location

param resourceIdentifier string = newGuid()

resource app 'Applications.Core/applications@2022-03-15-privatepreview' = {
  name: 'corerp-resources-dapr-pubsub-servicebus'
  location: location
  properties: {
    environment: environment
  }
}

resource publisher 'Applications.Core/containers@2022-03-15-privatepreview' = {
  name: 'sb-publisher'
  location: location
  properties: {
    application: app.id
    connections: {
      daprpubsub: {
        source: pubsub.id
      }
    }
    container: {
      image: magpieimage
      env: {
        BINDING_DAPRPUBSUB_NAME: pubsub.name
        BINDING_DAPRPUBSUB_TOPIC: pubsub.properties.topic
      }
      readinessProbe:{
        kind: 'httpGet'
        containerPort: 3000
        path: '/healthz'
      }
    }
    extensions: [
      {
        kind: 'daprSidecar'
        appId: 'sb-pubsub'
        appPort: 3000
      }
    ]
  }
}

resource pubsub 'Applications.Connector/daprPubSubBrokers@2022-03-15-privatepreview' = {
  name: 'sb-pubsub'
  location: location
  properties: {
    environment: environment
    application: app.id
    kind: 'pubsub.azure.servicebus'
    resource: namespace.id
  }
}

resource namespace 'Microsoft.ServiceBus/namespaces@2017-04-01' = {
  name: 'daprns-${resourceIdentifier}'
  location: location
  tags: {
    radiustest: 'corerp-resources-dapr-pubsub-servicebus'
  }
}
