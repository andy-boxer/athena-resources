const core = require('@actions/core');
const axios = require("axios");

const apiUrl = core.getInput('deployment_server');
const lidUid = core.getInput('lid_uid');
const namespace = core.getInput('namespace');
const url = apiUrl + '/lid/progress?namespace=' + namespace + '&uid=' + lidUid;

try {
  axios.get(url).then(response => {
    core.debug(response.data);
    core.setOutput("status", response.data.status );
    core.setOutput("uid", response.data.uid );
    core.setOutput("replicas", response.data.replicas );
  })
} catch (error) {
  core.setFailed(error.message);
}
