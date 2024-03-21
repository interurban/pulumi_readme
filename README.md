<p align="center">
  <a href="https://www.pulumi.com?utm_campaign=pulumi-pulumi-github-repo&utm_source=github.com&utm_medium=top-logo" title="Pulumi - Modern Infrastructure as Code - AWS Azure Kubernetes Containers Serverless">
    <img src="https://www.pulumi.com/images/logo/logo-on-white-box.svg?" width="350">
   </a>

  [![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com?utm_campaign=pulumi-pulumi-github-repo&utm_source=github.com&utm_medium=slack-badge)
  ![GitHub Discussions](https://img.shields.io/github/discussions/pulumi/pulumi)
  [![NPM version](https://badge.fury.io/js/%40pulumi%2Fpulumi.svg)](https://npmjs.com/package/@pulumi/pulumi)
  [![Python version](https://badge.fury.io/py/pulumi.svg)](https://pypi.org/project/pulumi)
  [![NuGet version](https://badge.fury.io/nu/pulumi.svg)](https://badge.fury.io/nu/pulumi)
  [![GoDoc](https://godoc.org/github.com/pulumi/pulumi?status.svg)](https://godoc.org/github.com/pulumi/pulumi)
  [![License](https://img.shields.io/github/license/pulumi/pulumi)](LICENSE)
  [![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/pulumi/pulumi)

  <a href="https://www.pulumi.com/docs/get-started/?utm_campaign=pulumi-pulumi-github-repo&utm_source=github.com&utm_medium=get-started-button" title="Get Started">
    <img src="https://www.pulumi.com/images/get-started.svg?" align="right" width="120">
  </a>
</p>

**Pulumi's Infrastructure as Code SDK** is the easiest way to build and deploy infrastructure, of any architecture and on any cloud, using programming languages that you already know and love. Code and ship infrastructure faster with your favorite languages and tools, and embed IaC anywhere with [Automation API](https://www.pulumi.com/docs/guides/automation-api/).

Simply write code in your favorite language and Pulumi automatically provisions and manages your resources on
[AWS](https://www.pulumi.com/docs/reference/clouds/aws/?utm_campaign=pulumi-pulumi-github-repo&utm_source=github.com&utm_medium=aws-reference-link),
[Azure](https://www.pulumi.com/docs/reference/clouds/azure/?utm_campaign=pulumi-pulumi-github-repo&utm_source=github.com&utm_medium=azure-reference-link),
[Google Cloud Platform](https://www.pulumi.com/docs/reference/clouds/gcp/?utm_campaign=pulumi-pulumi-github-repo&utm_source=github.com&utm_medium=gcp-reference-link), 
[Kubernetes](https://www.pulumi.com/docs/reference/clouds/kubernetes/?utm_campaign=pulumi-pulumi-github-repo&utm_source=github.com&utm_medium=kuberneters-reference-link), and [120+ providers](https://www.pulumi.com/registry/?utm_campaign=pulumi-pulumi-github-repo&utm_source=github.com&utm_medium=providers-reference-link) using an
[infrastructure-as-code](https://www.pulumi.com/what-is/what-is-infrastructure-as-code/) approach.
Skip the YAML, and use standard language features like loops, functions, classes,
and package management that you already know and love.

Pulumi is open source under the [Apache 2.0 license](https://github.com/pulumi/pulumi/blob/master/LICENSE), supports many languages and clouds, and is easy to extend.  This repo contains the `pulumi` CLI, language SDKs, and core Pulumi engine, and individual libraries are in their own repos.

## Table of contents

- :rocket: [Getting Started](#getting-started)
  - :blue_book: [Docs](#docs)
  - :school: [Learn](#learn)
  - :hammer_and_wrench: [Examples](#examples)
- :toolbox:	[Registry](#registry)
- :compass:	[Pulumi Roadmap](#pulumi-roadmap)
- :busts_in_silhouette: [Community](#community)
- :clap: [Contributing](#contributing)

## Getting Started
<img align="right" width="400" src="https://www.pulumi.com/images/docs/quickstart/console.png" />

[![Watch the video](/youtube_preview_image.png)](https://www.youtube.com/watch?v=6f8KF6UGN7g)

Visit our [Get Started with Pulumi](https://www.pulumi.com/docs/quickstart/?utm_campaign=pulumi-pulumi-github-repo&utm_source=github.com&utm_medium=getting-started-quickstart) guides to quickly deploy a simple application in AWS, Azure, Google Cloud, or Kubernetes using Pulumi.

Otherwise, the following steps demonstrate how to deploy your first Pulumi program, using AWS Serverless Lambdas, in minutes:

1. **Install**:

    To install the latest Pulumi release, run the following (see full
    [installation instructions](https://www.pulumi.com/docs/reference/install/?utm_campaign=pulumi-pulumi-github-repo&utm_source=github.com&utm_medium=getting-started-install) for additional installation options):

    ```bash
    $ curl -fsSL https://get.pulumi.com/ | sh
    ```

2. **Create a Project**:

    After installing, you can get started with the `pulumi new` command:

    ```bash
    $ mkdir pulumi-demo && cd pulumi-demo
    $ pulumi new hello-aws-javascript
    ```

    The `new` command offers templates for all languages and clouds.  Run it without an argument and it'll prompt
    you with available projects.  This command created an AWS Serverless Lambda project written in JavaScript.

3. **Deploy to the Cloud**:

    Run `pulumi up` to get your code to the cloud:

    ```bash
    $ pulumi up
    ```

    This makes all cloud resources needed to run your code.  Simply make edits to your project, and subsequent
    `pulumi up`s will compute the minimal diff to deploy your changes.

4. **Use Your Program**:

    Now that your code is deployed, you can interact with it.  In the above example, we can curl the endpoint:

    ```bash
    $ curl $(pulumi stack output url)
    ```

5. **Access the Logs**:

    If you're using containers or functions, Pulumi's unified logging command will show all of your logs:

    ```bash
    $ pulumi logs -f
    ```

6. **Destroy your Resources**:

    After you're done, you can remove all resources created by your program:

    ```bash
    $ pulumi destroy -y
    ```

### Documentation

Documentation including our Pulumi registry, with hundreds of packages and providers you can use to manage cloud services can be found on [Pulumi's docs site.](https://www.pulumi.com/docs/)

## Examples

Visit [pulumi/examples](https://github.com/pulumi/examples) for more than 100 supported examples spanning containers, serverless, and many other infrastructure as code use cases. Here are two of our favorites.

### Create three web servers in AWS

```typescript
const aws = require("@pulumi/aws");
const sg = new aws.ec2.SecurityGroup("web-sg", {
    ingress: [{ protocol: "tcp", fromPort: 80, toPort: 80, cidrBlocks: ["0.0.0.0/0"] }],
});
for (let i = 0; i < 3; i++) {
    new aws.ec2.Instance(`web-${i}`, {
        ami: "ami-7172b611",
        instanceType: "t2.micro",
        vpcSecurityGroupIds: [sg.id],
        userData: `#!/bin/bash
            echo "Hello, World!" > index.html
            nohup python -m SimpleHTTPServer 80 &`,
    });
}
```

### Create a simple serverless timer that archives Hacker News every day at 8:30AM in AWS

```typescript
const aws = require("@pulumi/aws");

const snapshots = new aws.dynamodb.Table("snapshots", {
    attributes: [{ name: "id", type: "S", }],
    hashKey: "id", billingMode: "PAY_PER_REQUEST",
});

aws.cloudwatch.onSchedule("daily-yc-snapshot", "cron(30 8 * * ? *)", () => {
    require("https").get("https://news.ycombinator.com", res => {
        let content = "";
        res.setEncoding("utf8");
        res.on("data", chunk => content += chunk);
        res.on("end", () => new aws.sdk.DynamoDB.DocumentClient().put({
            TableName: snapshots.name.get(),
            Item: { date: Date.now(), content },
        }).promise());
    }).end();
});
```

### [Learn Pulumi](https://www.pulumi.com/learn/): Follow Pulumi learning pathways for a step-by-step guide for using Pulumi for the first time.

## Pulumi Roadmap

Review the planned work for the upcoming quarter and a selected backlog of issues that are on our mind but not yet scheduled on the [Pulumi Roadmap.](https://github.com/orgs/pulumi/projects/44)

## Community

- Join us in the [Pulumi Community Slack](https://slack.pulumi.com/?utm_campaign=pulumi-pulumi-github-repo&utm_source=github.com&utm_medium=welcome-slack) to connect with our community and engineering team and ask questions. All conversations and questions are welcome.
- Send us a tweet via [(]@PulumiCorp](https://twitter.com/PulumiCorp)
- Watch videos and workshops on [Pulumi TV](https://www.youtube.com/pulumitv)

### Discussions

You can also visit our [GitHub Discussions](https://github.com/pulumi/pulumi/discussions) to ask questions or share what you're building with Pulumi.

## Languages and Providers

|    | Language | Status | Runtime | Versions |
| -- | -------- | ------ | ------- | -------- |
| <img src="https://www.pulumi.com/logos/tech/logo-js.png" height=38 />     | [JavaScript](https://www.pulumi.com/docs/intro/languages/javascript/) | Stable  | Node.js | [Current, Active and Maintenance LTS versions](https://nodejs.org/en/about/previous-releases)  |
| <img src="https://www.pulumi.com/logos/tech/logo-ts.png" height=38 />     | [TypeScript](https://www.pulumi.com/docs/intro/languages/javascript/) | Stable  | Node.js | [Current, Active and Maintenance LTS versions](https://nodejs.org/en/about/previous-releases)  |
| <img src="https://www.pulumi.com/logos/tech/logo-python.svg" height=38 /> | [Python](https://www.pulumi.com/docs/intro/languages/python/)     | Stable  | Python | [Supported versions](https://devguide.python.org/versions/#versions) |
| <img src="https://www.pulumi.com/logos/tech/logo-golang.png" height=38 /> | [Go](https://www.pulumi.com/docs/intro/languages/go/)             | Stable  | Go | [Supported versions](https://go.dev/doc/devel/release#policy) |
| <img src="https://www.pulumi.com/logos/tech/dotnet.svg" height=38 />      | [.NET (C#/F#/VB.NET)](https://www.pulumi.com/docs/intro/languages/dotnet/)     | Stable  | .NET | [Supported versions](https://dotnet.microsoft.com/en-us/platform/support/policy/dotnet-core#lifecycle)  |
| <img src="https://www.pulumi.com/logos/tech/java.svg" height=38 />      | [Java](https://www.pulumi.com/docs/intro/languages/java/)     | Public Preview  | JDK | 11+  |
| <img src="https://www.pulumi.com/logos/tech/yaml.svg" height=38 />      | [YAML](https://www.pulumi.com/docs/intro/languages/yaml/)     | Stable  | n/a  | n/a  |

Visit the [Pulumi Registry](https://www.pulumi.com/registry/) for the full list of supported cloud and infrastructure providers.

## EOL Releases

The Pulumi CLI v1 and v2 are no longer supported. If you are not yet running v3, please consider migrating to v3 to continue getting the latest and greatest Pulumi has to offer! :muscle:

* To migrate from v2 to v3, please see our [v3 Migration Guide](https://www.pulumi.com/docs/install/migrating-3.0/).

## Contributing

Contributions to Pulumi are welcome. Read our [contributing guide](https://github.com/pulumi/pulumi/blob/master/CONTRIBUTING.md) for information on building Pulumi from source or contributing improvements.

Looking for a first issue to tackle?

- We tag issues with [![Good First Issue](linktbd) for people new to our codebase or OSS contributions in general.