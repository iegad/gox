namespace Sphinx.Forms
{
    partial class CerberusFm
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
            this.txtURL = new System.Windows.Forms.TextBox();
            this.gbCerberusURL = new System.Windows.Forms.GroupBox();
            this.tvNodes = new System.Windows.Forms.TreeView();
            this.btnGetInfo = new System.Windows.Forms.Button();
            this.groupBox1 = new System.Windows.Forms.GroupBox();
            this.lblHttp = new System.Windows.Forms.Label();
            this.gbCerberusURL.SuspendLayout();
            this.SuspendLayout();
            // 
            // txtURL
            // 
            this.txtURL.Anchor = ((System.Windows.Forms.AnchorStyles)(((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.txtURL.Location = new System.Drawing.Point(100, 32);
            this.txtURL.Name = "txtURL";
            this.txtURL.Size = new System.Drawing.Size(370, 33);
            this.txtURL.TabIndex = 1;
            // 
            // gbCerberusURL
            // 
            this.gbCerberusURL.Anchor = ((System.Windows.Forms.AnchorStyles)(((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left)));
            this.gbCerberusURL.Controls.Add(this.lblHttp);
            this.gbCerberusURL.Controls.Add(this.btnGetInfo);
            this.gbCerberusURL.Controls.Add(this.tvNodes);
            this.gbCerberusURL.Controls.Add(this.txtURL);
            this.gbCerberusURL.Location = new System.Drawing.Point(12, 12);
            this.gbCerberusURL.Name = "gbCerberusURL";
            this.gbCerberusURL.Size = new System.Drawing.Size(611, 999);
            this.gbCerberusURL.TabIndex = 2;
            this.gbCerberusURL.TabStop = false;
            this.gbCerberusURL.Text = "Cerberus URL";
            // 
            // tvNodes
            // 
            this.tvNodes.Anchor = ((System.Windows.Forms.AnchorStyles)((((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.tvNodes.Location = new System.Drawing.Point(6, 71);
            this.tvNodes.Name = "tvNodes";
            this.tvNodes.Size = new System.Drawing.Size(599, 922);
            this.tvNodes.TabIndex = 3;
            // 
            // btnGetInfo
            // 
            this.btnGetInfo.Anchor = ((System.Windows.Forms.AnchorStyles)(((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.btnGetInfo.Location = new System.Drawing.Point(476, 32);
            this.btnGetInfo.Name = "btnGetInfo";
            this.btnGetInfo.Size = new System.Drawing.Size(129, 33);
            this.btnGetInfo.TabIndex = 4;
            this.btnGetInfo.Text = "获取信息";
            this.btnGetInfo.UseVisualStyleBackColor = true;
            // 
            // groupBox1
            // 
            this.groupBox1.Anchor = ((System.Windows.Forms.AnchorStyles)(((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.groupBox1.Location = new System.Drawing.Point(629, 12);
            this.groupBox1.Name = "groupBox1";
            this.groupBox1.Size = new System.Drawing.Size(1100, 359);
            this.groupBox1.TabIndex = 3;
            this.groupBox1.TabStop = false;
            this.groupBox1.Text = "groupBox1";
            // 
            // lblHttp
            // 
            this.lblHttp.Anchor = ((System.Windows.Forms.AnchorStyles)(((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.lblHttp.AutoSize = true;
            this.lblHttp.Location = new System.Drawing.Point(7, 37);
            this.lblHttp.Name = "lblHttp";
            this.lblHttp.Size = new System.Drawing.Size(87, 22);
            this.lblHttp.TabIndex = 5;
            this.lblHttp.Text = "http://";
            // 
            // CerberusFm
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(11F, 22F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(1741, 1023);
            this.Controls.Add(this.groupBox1);
            this.Controls.Add(this.gbCerberusURL);
            this.Font = new System.Drawing.Font("宋体", 11F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            this.Margin = new System.Windows.Forms.Padding(4, 4, 4, 4);
            this.Name = "CerberusFm";
            this.Text = "CerberusFm";
            this.FormClosed += new System.Windows.Forms.FormClosedEventHandler(this.CerberusFm_FormClosed);
            this.gbCerberusURL.ResumeLayout(false);
            this.gbCerberusURL.PerformLayout();
            this.ResumeLayout(false);

        }

        #endregion

        private System.Windows.Forms.TextBox txtURL;
        private System.Windows.Forms.GroupBox gbCerberusURL;
        private System.Windows.Forms.TreeView tvNodes;
        private System.Windows.Forms.Label lblHttp;
        private System.Windows.Forms.Button btnGetInfo;
        private System.Windows.Forms.GroupBox groupBox1;
    }
}