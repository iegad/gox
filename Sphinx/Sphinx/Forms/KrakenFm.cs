using Newtonsoft.Json.Linq;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace Sphinx.Forms
{
    public partial class KrakenFm : Form
    {
        private static KrakenFm instance_;
        private string ManagerHost_;
        private ListViewItem SelectedKraken_;
        private long Idempontent_;

        public static KrakenFm Instance {
            get {
                if (instance_ == null) {
                    instance_ = new KrakenFm();
                }

                return instance_;
            }
        }


        KrakenFm()
        {
            InitializeComponent();
        }

        private void KrakenFm_FormClosed(object sender, FormClosedEventArgs e)
        {
            Close();
        }

        public new void Close()
        {
            Dispose();
        }

        private void KrakenFm_Load(object sender, EventArgs e)
        {
            var lvi = new ListViewItem("1");
            lvi.SubItems.Add("ae906ff9-552b-4f63-b807-95a1195deddf");
            lvi.SubItems.Add("127.0.0.1:9090");
            lvi.SubItems.Add("");
            lvi.SubItems.Add("");
            lvKrakens.Items.Add(lvi);

            cmbCGIs.SelectedIndex = 0;
        }

        private void cmbCGIs_SelectedIndexChanged(object sender, EventArgs e)
        {
            if (lvKrakens.SelectedItems.Count == 0) {
                gbDetails.Text = "URL";
                return;
            }

            int userID = 0;
            string token = "";

            if (SelectedKraken_ != null) {
                string tmp = SelectedKraken_.SubItems[3].Text.Trim();
                if (tmp.Length > 0) {
                    userID = Convert.ToInt32(tmp);
                }

                token = SelectedKraken_.SubItems[4].Text.Trim();
            }

            if (cmbCGIs.SelectedIndex > 0 && (userID == 0 || token.Length != 16)) {
                return;
            }

            switch(cmbCGIs.SelectedIndex)
            {
                case 0:
                    gbDetails.Text = string.Format("http://{0}/login", ManagerHost_);
                    rtxtReq.Text = "{\"user_name\": \"root\", \"password\": \"e807f1fcf82d132f9bb018ca6738a19f\"}";
                    rtxtRsp.Text = "";
                    break;

                case 1:
                    gbDetails.Text = string.Format("http://{0}/front/shutdown", ManagerHost_);
                    rtxtReq.Text = string.Format("{{\"user_id\": {0}, \"token\": \"{1}\", \"idempotent\": {2}}}", userID, token, ++Idempontent_);
                    rtxtRsp.Text = "";
                    break;

                case 2:
                    gbDetails.Text = string.Format("http://{0}/front/run", ManagerHost_);
                    rtxtReq.Text = string.Format("{{\"user_id\": {0}, \"token\": \"{1}\", \"idempotent\": {2}}}", userID, token, ++Idempontent_);
                    rtxtRsp.Text = "";
                    break;

                case 3:
                    gbDetails.Text = string.Format("http://{0}/front/info", ManagerHost_);
                    rtxtReq.Text = string.Format("{{\"user_id\": {0}, \"token\": \"{1}\", \"idempotent\": {2}}}", userID, token, ++Idempontent_);
                    rtxtRsp.Text = "";
                    break;

                case 4:
                    gbDetails.Text = string.Format("http://{0}/front/kick_session", ManagerHost_);
                    break;

                case 5:
                    gbDetails.Text = string.Format("http://{0}/backend/shutdown", ManagerHost_);
                    rtxtReq.Text = string.Format("{{\"user_id\": {0}, \"token\": \"{1}\", \"idempotent\": {2}}}", userID, token, ++Idempontent_);
                    rtxtRsp.Text = "";
                    break;

                case 6:
                    gbDetails.Text = string.Format("http://{0}/backend/run", ManagerHost_);
                    rtxtReq.Text = string.Format("{{\"user_id\": {0}, \"token\": \"{1}\", \"idempotent\": {2}}}", userID, token, ++Idempontent_);
                    rtxtRsp.Text = "";
                    break;

                case 7:
                    gbDetails.Text = string.Format("http://{0}/backend/info", ManagerHost_);
                    rtxtReq.Text = string.Format("{{\"user_id\": {0}, \"token\": \"{1}\", \"idempotent\": {2}}}", userID, token, ++Idempontent_);
                    rtxtRsp.Text = "";
                    break;

                default:
                    gbDetails.Text = "URL"; 
                    break;
            }
        }

        private void lvKrakens_SelectedIndexChanged(object sender, EventArgs e)
        {
            if (lvKrakens.SelectedItems.Count == 0) {
                return;
            }

            SelectedKraken_ = lvKrakens.SelectedItems[0];
            ManagerHost_ = SelectedKraken_.SubItems[2].Text;
            cmbCGIs_SelectedIndexChanged(null, null);
        }

        private void btnSend_Click(object sender, EventArgs e)
        {
            if (SelectedKraken_ == null) {
                return;
            }

            string url = gbDetails.Text;
            string req = rtxtReq.Text;
            string rsp = Utils.HttpPost(url, req);
            rtxtRsp.Text = rsp;

            if (rsp.Length > 0) {
                JObject tmp = JObject.Parse(rsp);
                int code = tmp["code"].Value<int>();
                if (code == 0) {
                    switch (cmbCGIs.SelectedIndex)
                    {
                        case 0:
                            JObject result = JObject.Parse(tmp["data"].ToString());
                            SelectedKraken_.SubItems[3].Text = result["user_id"].ToString();
                            SelectedKraken_.SubItems[4].Text = result["token"].ToString();
                            break;

                        case 1:
                        case 2:
                        case 3:
                        case 4:
                        case 5:
                        case 6:
                        case 7:
                            break;
                    }
                }
            }
        }
    }
}
