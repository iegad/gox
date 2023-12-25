namespace Sphinx
{
    partial class MainFm
    {
        /// <summary>
        /// 必需的设计器变量。
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// 清理所有正在使用的资源。
        /// </summary>
        /// <param name="disposing">如果应释放托管资源，为 true；否则为 false。</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows 窗体设计器生成的代码

        /// <summary>
        /// 设计器支持所需的方法 - 不要修改
        /// 使用代码编辑器修改此方法的内容。
        /// </summary>
        private void InitializeComponent()
        {
            this.topMenu = new System.Windows.Forms.MenuStrip();
            this.Kraken = new System.Windows.Forms.ToolStripMenuItem();
            this.miKrakenCGI = new System.Windows.Forms.ToolStripMenuItem();
            this.topMenu.SuspendLayout();
            this.SuspendLayout();
            // 
            // topMenu
            // 
            this.topMenu.GripMargin = new System.Windows.Forms.Padding(2, 2, 0, 2);
            this.topMenu.ImageScalingSize = new System.Drawing.Size(24, 24);
            this.topMenu.Items.AddRange(new System.Windows.Forms.ToolStripItem[] {
            this.Kraken});
            this.topMenu.Location = new System.Drawing.Point(0, 0);
            this.topMenu.Name = "topMenu";
            this.topMenu.Size = new System.Drawing.Size(1901, 32);
            this.topMenu.TabIndex = 1;
            this.topMenu.Text = "menuStrip1";
            // 
            // Kraken
            // 
            this.Kraken.DropDownItems.AddRange(new System.Windows.Forms.ToolStripItem[] {
            this.miKrakenCGI});
            this.Kraken.Name = "Kraken";
            this.Kraken.Size = new System.Drawing.Size(62, 28);
            this.Kraken.Text = "网关";
            // 
            // miKrakenCGI
            // 
            this.miKrakenCGI.Name = "miKrakenCGI";
            this.miKrakenCGI.Size = new System.Drawing.Size(270, 34);
            this.miKrakenCGI.Text = "网关CGI测试";
            this.miKrakenCGI.Click += new System.EventHandler(this.miKrakenCGI_Click);
            // 
            // MainFm
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(9F, 18F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(1901, 1107);
            this.Controls.Add(this.topMenu);
            this.IsMdiContainer = true;
            this.MainMenuStrip = this.topMenu;
            this.Name = "MainFm";
            this.Text = "Sphinx";
            this.topMenu.ResumeLayout(false);
            this.topMenu.PerformLayout();
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.MenuStrip topMenu;
        private System.Windows.Forms.ToolStripMenuItem Kraken;
        private System.Windows.Forms.ToolStripMenuItem miKrakenCGI;
    }
}

