namespace Sphinx.Forms
{
    partial class KrakenFm
    {
        /// <summary>
        /// Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            Hide();
        }

        #region Windows Form Designer generated code

        /// <summary>
        /// Required method for Designer support - do not modify
        /// the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            this.lvKrakens = new System.Windows.Forms.ListView();
            this.SeqNo = ((System.Windows.Forms.ColumnHeader)(new System.Windows.Forms.ColumnHeader()));
            this.NodeCode = ((System.Windows.Forms.ColumnHeader)(new System.Windows.Forms.ColumnHeader()));
            this.ManagerHost = ((System.Windows.Forms.ColumnHeader)(new System.Windows.Forms.ColumnHeader()));
            this.UserID = ((System.Windows.Forms.ColumnHeader)(new System.Windows.Forms.ColumnHeader()));
            this.Token = ((System.Windows.Forms.ColumnHeader)(new System.Windows.Forms.ColumnHeader()));
            this.gbCGIs = new System.Windows.Forms.GroupBox();
            this.btnSend = new System.Windows.Forms.Button();
            this.gbDetails = new System.Windows.Forms.GroupBox();
            this.rtxtRsp = new System.Windows.Forms.RichTextBox();
            this.rtxtReq = new System.Windows.Forms.RichTextBox();
            this.lblRsp = new System.Windows.Forms.Label();
            this.lblReq = new System.Windows.Forms.Label();
            this.cmbCGIs = new System.Windows.Forms.ComboBox();
            this.gbCGIs.SuspendLayout();
            this.gbDetails.SuspendLayout();
            this.SuspendLayout();
            // 
            // lvKrakens
            // 
            this.lvKrakens.Anchor = ((System.Windows.Forms.AnchorStyles)(((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.lvKrakens.Columns.AddRange(new System.Windows.Forms.ColumnHeader[] {
            this.SeqNo,
            this.NodeCode,
            this.ManagerHost,
            this.UserID,
            this.Token});
            this.lvKrakens.Font = new System.Drawing.Font("宋体", 12F);
            this.lvKrakens.FullRowSelect = true;
            this.lvKrakens.HideSelection = false;
            this.lvKrakens.Location = new System.Drawing.Point(12, 12);
            this.lvKrakens.Name = "lvKrakens";
            this.lvKrakens.Size = new System.Drawing.Size(1700, 445);
            this.lvKrakens.TabIndex = 0;
            this.lvKrakens.UseCompatibleStateImageBehavior = false;
            this.lvKrakens.View = System.Windows.Forms.View.Details;
            this.lvKrakens.SelectedIndexChanged += new System.EventHandler(this.lvKrakens_SelectedIndexChanged);
            // 
            // SeqNo
            // 
            this.SeqNo.Text = "序号";
            this.SeqNo.Width = 100;
            // 
            // NodeCode
            // 
            this.NodeCode.Text = "节点编码";
            this.NodeCode.Width = 420;
            // 
            // ManagerHost
            // 
            this.ManagerHost.Text = "管理服务";
            this.ManagerHost.Width = 240;
            // 
            // UserID
            // 
            this.UserID.Text = "UserID";
            this.UserID.Width = 115;
            // 
            // Token
            // 
            this.Token.Text = "Token";
            this.Token.Width = 231;
            // 
            // gbCGIs
            // 
            this.gbCGIs.Anchor = ((System.Windows.Forms.AnchorStyles)((((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.gbCGIs.Controls.Add(this.btnSend);
            this.gbCGIs.Controls.Add(this.gbDetails);
            this.gbCGIs.Controls.Add(this.cmbCGIs);
            this.gbCGIs.Font = new System.Drawing.Font("宋体", 12F);
            this.gbCGIs.Location = new System.Drawing.Point(12, 463);
            this.gbCGIs.Name = "gbCGIs";
            this.gbCGIs.Size = new System.Drawing.Size(1700, 531);
            this.gbCGIs.TabIndex = 1;
            this.gbCGIs.TabStop = false;
            this.gbCGIs.Text = "CGIs";
            // 
            // btnSend
            // 
            this.btnSend.Location = new System.Drawing.Point(6, 82);
            this.btnSend.Name = "btnSend";
            this.btnSend.Size = new System.Drawing.Size(139, 44);
            this.btnSend.TabIndex = 3;
            this.btnSend.Text = "发送";
            this.btnSend.UseVisualStyleBackColor = true;
            this.btnSend.Click += new System.EventHandler(this.btnSend_Click);
            // 
            // gbDetails
            // 
            this.gbDetails.Anchor = ((System.Windows.Forms.AnchorStyles)((((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.gbDetails.Controls.Add(this.rtxtRsp);
            this.gbDetails.Controls.Add(this.rtxtReq);
            this.gbDetails.Controls.Add(this.lblRsp);
            this.gbDetails.Controls.Add(this.lblReq);
            this.gbDetails.Location = new System.Drawing.Point(426, 34);
            this.gbDetails.Name = "gbDetails";
            this.gbDetails.Size = new System.Drawing.Size(1268, 491);
            this.gbDetails.TabIndex = 2;
            this.gbDetails.TabStop = false;
            this.gbDetails.Text = "URL";
            // 
            // rtxtRsp
            // 
            this.rtxtRsp.Anchor = ((System.Windows.Forms.AnchorStyles)((((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.rtxtRsp.Location = new System.Drawing.Point(10, 285);
            this.rtxtRsp.Name = "rtxtRsp";
            this.rtxtRsp.Size = new System.Drawing.Size(1252, 200);
            this.rtxtRsp.TabIndex = 3;
            this.rtxtRsp.Text = "";
            // 
            // rtxtReq
            // 
            this.rtxtReq.Anchor = ((System.Windows.Forms.AnchorStyles)(((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.rtxtReq.Location = new System.Drawing.Point(10, 55);
            this.rtxtReq.Name = "rtxtReq";
            this.rtxtReq.Size = new System.Drawing.Size(1252, 200);
            this.rtxtReq.TabIndex = 2;
            this.rtxtReq.Text = "";
            // 
            // lblRsp
            // 
            this.lblRsp.Anchor = ((System.Windows.Forms.AnchorStyles)(((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.lblRsp.AutoSize = true;
            this.lblRsp.Location = new System.Drawing.Point(6, 258);
            this.lblRsp.Name = "lblRsp";
            this.lblRsp.Size = new System.Drawing.Size(106, 24);
            this.lblRsp.TabIndex = 1;
            this.lblRsp.Text = "Response";
            // 
            // lblReq
            // 
            this.lblReq.AutoSize = true;
            this.lblReq.Location = new System.Drawing.Point(6, 31);
            this.lblReq.Name = "lblReq";
            this.lblReq.Size = new System.Drawing.Size(94, 24);
            this.lblReq.TabIndex = 0;
            this.lblReq.Text = "Request";
            // 
            // cmbCGIs
            // 
            this.cmbCGIs.DropDownStyle = System.Windows.Forms.ComboBoxStyle.DropDownList;
            this.cmbCGIs.Font = new System.Drawing.Font("宋体", 12F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            this.cmbCGIs.FormattingEnabled = true;
            this.cmbCGIs.Items.AddRange(new object[] {
            "Login(用户登录)",
            "Front/Shutdown(关闭前端服务)",
            "Front/Run(启动前端服务)",
            "Front/Info(前端服务信息)",
            "Front/Kick_Session(踢除会话)",
            "/Backend/Shutdown(关闭后端服务)",
            "/Backend/Run(启动后端服务)",
            "/Backend/Info(后端服务信息)"});
            this.cmbCGIs.Location = new System.Drawing.Point(6, 34);
            this.cmbCGIs.Name = "cmbCGIs";
            this.cmbCGIs.Size = new System.Drawing.Size(414, 32);
            this.cmbCGIs.TabIndex = 0;
            this.cmbCGIs.SelectedIndexChanged += new System.EventHandler(this.cmbCGIs_SelectedIndexChanged);
            // 
            // KrakenFm
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(9F, 18F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(1724, 1006);
            this.Controls.Add(this.gbCGIs);
            this.Controls.Add(this.lvKrakens);
            this.Name = "KrakenFm";
            this.Text = "Kraken";
            this.FormClosed += new System.Windows.Forms.FormClosedEventHandler(this.KrakenFm_FormClosed);
            this.Load += new System.EventHandler(this.KrakenFm_Load);
            this.gbCGIs.ResumeLayout(false);
            this.gbDetails.ResumeLayout(false);
            this.gbDetails.PerformLayout();
            this.ResumeLayout(false);

        }

        #endregion

        private System.Windows.Forms.ListView lvKrakens;
        private System.Windows.Forms.ColumnHeader SeqNo;
        private System.Windows.Forms.ColumnHeader NodeCode;
        private System.Windows.Forms.ColumnHeader ManagerHost;
        private System.Windows.Forms.GroupBox gbCGIs;
        private System.Windows.Forms.ComboBox cmbCGIs;
        private System.Windows.Forms.GroupBox gbDetails;
        private System.Windows.Forms.Label lblReq;
        private System.Windows.Forms.Label lblRsp;
        private System.Windows.Forms.RichTextBox rtxtRsp;
        private System.Windows.Forms.RichTextBox rtxtReq;
        private System.Windows.Forms.Button btnSend;
        private System.Windows.Forms.ColumnHeader Token;
        private System.Windows.Forms.ColumnHeader UserID;
    }
}