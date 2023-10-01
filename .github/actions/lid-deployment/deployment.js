const core = require('@actions/core');
const github = require('@actions/github');
const axios = require("axios");

const apiUrl = core.getInput('deployment_server');
const imageRef = core.getInput('image_ref');
const namespace = core.getInput('namespace');
const deploymentKind = core.getInput('deployment_kind');

const lidRequest = {
  imageRef: imageRef,
  namespace: namespace,
  deploymentKind: deploymentKind,
  deployer: github.context.actor,
};

try {
  axios.post(apiUrl + "/lid/deploy", lidRequest).then(response => {
    core.setOutput("status", response.data.status);
    core.setOutput("uid", response.data.lidUid);
  })
} catch (error) {
  core.setFailed(error.message);
}
