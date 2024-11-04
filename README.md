# Amazon SageMaker Operators for Kubernetes

![](https://i.imgur.com/waxVImv.png)
### [View all Roadmaps](https://github.com/nholuongut/all-roadmaps) &nbsp;&middot;&nbsp; [Best Practices](https://github.com/nholuongut/all-roadmaps/blob/main/public/best-practices/) &nbsp;&middot;&nbsp; [Questions](https://www.linkedin.com/in/nholuong/)
<br/>

## Introduction
Amazon SageMaker Operators for Kubernetes are operators that can be used to train machine learning models, optimize hyperparameters for a given model, run batch transform jobs over existing models, and set up inference endpoints. With these operators, users can manage their jobs in Amazon SageMaker from their Kubernetes cluster in Amazon Elastic Kubernetes Service [EKS](http://aws.amazon.com/eks).

## Usage

First, you must [install the operators](https://sagemaker.readthedocs.io/en/stable/amazon_sagemaker_operators_for_kubernetes.html). After installation is complete, create a TrainingJob YAML specification by following one of the samples, like [samples/xgboost-mnist-trainingjob.yaml](./samples/xgboost-mnist-trainingjob.yaml). Then, use `kubectl` to create and monitor the progress of your job:

```bash
$ kubectl apply -f xgboost-mnist-trainingjob.yaml
trainingjob.sagemaker.aws.amazon.com/xgboost-mnist created

$ kubectl get trainingjob
NAME            STATUS       SECONDARY-STATUS   CREATION-TIME          SAGEMAKER-JOB-NAME
xgboost-mnist   InProgress   Starting           2019-11-26T23:38:11Z   xgboost-mnist-cf1e16fb10a511eaaa450a350733ba06
```

Once the job starts training, you can use a `kubectl` plugin to stream training logs:

```bash
$ kubectl get trainingjob
NAME            STATUS       SECONDARY-STATUS   CREATION-TIME          SAGEMAKER-JOB-NAME
xgboost-mnist   InProgress   Training           2019-11-26T23:38:11Z   xgboost-mnist-cf1e16fb10a511eaaa450a350733ba06

$ kubectl smlogs trainingjob xgboost-mnist | head -n 5
"xgboost-mnist" has SageMaker TrainingJobName "xgboost-mnist-cf1e16fb10a511eaaa450a350733ba06" in region "us-east-2", status "InProgress" and secondary status "Training"
xgboost-mnist-cf1e16fb10a511eaaa450a350733ba06/algo-1-1574811611 2019-11-26 15:41:13.449 -0800 PST Arguments: train
xgboost-mnist-cf1e16fb10a511eaaa450a350733ba06/algo-1-1574811611 2019-11-26 15:41:13.449 -0800 PST [2019-11-26:23:41:10:INFO] Running standalone xgboost training.
xgboost-mnist-cf1e16fb10a511eaaa450a350733ba06/algo-1-1574811611 2019-11-26 15:41:13.45 -0800 PST [2019-11-26:23:41:10:INFO] File size need to be processed in the node: 1122.95mb. Available memory size in the node: 8501.08mb
xgboost-mnist-cf1e16fb10a511eaaa450a350733ba06/algo-1-1574811611 2019-11-26 15:41:13.45 -0800 PST [2019-11-26:23:41:10:INFO] Determined delimiter of CSV input is ','
xgboost-mnist-cf1e16fb10a511eaaa450a350733ba06/algo-1-1574811611 2019-11-26 15:41:13.45 -0800 PST [23:41:10] S3DistributionType set as FullyReplicated
```

The Amazon SageMaker Operators for Kubernetes enable management of SageMaker TrainingJobs, HyperParameterTuningJobs, BatchTransformJobs and HostingDeployments (Endpoints). Create and monitor them using the same `kubectl` tool as above.

To install the operators onto your Kubernetes cluster, follow our [User Guide](https://sagemaker.readthedocs.io/en/stable/amazon_sagemaker_operators_for_kubernetes.html).

### YAML Examples

To make a YAML spec, follow one of the below examples as a guide. Replace values like RoleARN, S3 input buckets and S3 output buckets with values that correspond to your account.

* [BatchTransformJob](./samples/xgboost-mnist-batchtransform.yaml)
* [HostingDeployment (Endpoint)](./samples/xgboost-mnist-hostingdeployment.yaml)
* [HyperParameterTuningJob](./samples/xgboost-mnist-hpo.yaml)
* [TrainingJob](./samples/xgboost-mnist-trainingjob.yaml)

## Releases

Amazon SageMaker Operator for Kubernetes adheres to the [SemVer](https://semver.org/) specification. Each release updates the major version tag (eg. `vX`), a major/minor version tag (eg. `vX.Y`) and a major/minor/patch version tag (eg. `vX.Y.Z`), as well as new versions of the `smlogs` binary with URLs of the same versioning formats. To see a full list of all releases, refer to our [Github releases page](https://github.com/nholuongut/amazon-sagemaker-operator-for-k8s/releases).

We also maintain a `latest` tag, which is updated to stay in line with the `master` branch. We **do not** recommend installing this on any production cluster, as any new major versions updated on the `master` branch will introduce breaking changes.

## Contributing
`amazon-sagemaker-operator-for-k8s` is an open source project. See [CONTRIBUTING](https://github.com/nholuongut/amazon-sagemaker-operator-for-k8s/blob/master/CONTRIBUTING.md) for details.

LABEL maintainer="Nho Luong <luongutnho@hotmail.com>"

# 🚀 I'm are always open to your feedback.  Please contact as bellow information:
### [Contact ]
* [Name: Nho Luong]
* [Skype](luongutnho_skype)
* [Github](https://github.com/nholuongut/)
* [Linkedin](https://www.linkedin.com/in/nholuong/)
* [Email Address](luongutnho@hotmail.com)
* [PayPal.me](https://www.paypal.com/paypalme/nholuongut)

![](https://i.imgur.com/waxVImv.png)
![](Donate.png)
[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/nholuong)

# License