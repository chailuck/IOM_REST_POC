create table CST_SUBR_PCN_00 (
	subr_row_id char(15) not null,
	cust_numb integer  not null,
	subr_numb char(12)  not null,
	blpd_code char(2)  not null,
	subr_type char(1)  ,
	neo_subr_type char(1)  ,
	telp_type char(1)  not null,
	titl nvarchar(40,10) ,
	frst_name nvarchar(80,40) ,
	last_name nvarchar(80,40) ,
	subr_stts char(1)  not null,
	from_optr_code char(3)  ,
	mnpi_stts char(1)  ,
	mnpo_stts char(1)  ,
	frst_pkgp_code char(8)  ,
	pkgp_strt_date date  ,
	new_ppty_code char(3)  ,
	dpst_txtp_code char(4)  ,
	conx_txtp_code char(4)  ,
	telp_allc_type char(1)  ,
	hrdw_numb char(20)  ,
	card_numb char(19)  not null,
	sms_lang char(1) default 'T',
	ivr_lang char(1) default 'T',
	ussd_lang char(1) default 'T',
	cc_lang char(2) default '01',
	deal_numb char(8)  not null,
	slmn_code char(5)  ,
	agmt_chck_flag char(1)  ,
	agmt_chck_date date  ,
	agmt_chck_by char(12)  ,
	frst_swon_dttm datetime year to second  ,
	brth_swon_dttm datetime year to second  ,
	swon_dttm datetime year to second  not null,
	swon_by char(12)  not null,
	swon_resn_code char(8)  not null,
	swon_area_code char(1)  ,
	swof_dttm datetime year to second  ,
	swof_by char(12)  ,
	swof_resn_code char(8)  ,
	rcnx_resn_code char(8)  ,
	leas_code char(5)  ,
	leas_strt_date date  ,
	leas_expr_date date  ,
	leas_acex_date date  ,
	pswd_flag char(1)  ,
	pswd char(4)  ,
	memb_id char(20) , 
	remk varchar(255)  ,
	id_type char(2)  ,
	id_numb char(20)  ,
	gndr char(1)  ,
	occp_code char(3)  ,
	marl_stts char(1)  ,
	salr_levl char(1)  ,
	educ_levl smallint  ,
	emal_addr varchar(40)  ,
	vrfy_emal_dttm datetime year to second  ,
	wait_emal_addr varchar(40)  ,
	wait_emal_dttm datetime year to second  ,
	rcnx_paid_flag char(1)  ,
	stop_bill_flag char(1)  ,
	rfnd_dttm datetime year to second  ,
	old_cust_numb integer  ,
	old_subr_numb char(12)  ,
	new_cust_numb integer  ,
	new_subr_numb char(12)  ,
	go_inter_flag char(1) default 'N' ,
	chwr_cont_aou char(1) default '0' ,
	last_chng_dttm datetime year to second  not null,
	last_chng_by char(12)  not null,

	primary key (subr_row_id)
		constraint cstp_subr_pcn_00_01
);

--==============================================================
-- Index: cstp_subr_pcn_00_01
--==============================================================
create unique index cstp_subr_pcn_00_01 on cst_subr_pcn_00 (
	subr_row_id  ASC
);


--==============================================================
-- Index: cstn_subr_pcn_00_02
--==============================================================
create unique index cstn_subr_pcn_00_02 on cst_subr_pcn_00 (
	cust_numb, subr_numb ASC
);

--==============================================================
-- Index: cstn_subr_pcn_00_03
--==============================================================
create index cstn_subr_pcn_00_03 on cst_subr_pcn_00 (
	subr_stts ASC
);


--==============================================================
-- Index: cstn_subr_pcn_00_04
--==============================================================
create index cstn_subr_pcn_00_04 on cst_subr_pcn_00 (
	swon_dttm ASC
);

--==============================================================
-- Index: cstn_subr_pcn_00_05
--==============================================================
create index cstn_subr_pcn_00_05 on cst_subr_pcn_00 (
	swof_dttm ASC
);

--==============================================================
-- Index: cstn_subr_pcn_00_06
--==============================================================
create index cstn_subr_pcn_00_06 on cst_subr_pcn_00 (
	old_cust_numb, old_subr_numb ASC
);

--==============================================================
-- Index: cstn_subr_pcn_00_07
--==============================================================
create index cstn_subr_pcn_00_07 on cst_subr_pcn_00 (
	telp_type ASC
);

--==============================================================
-- Index: cstn_subr_pcn_00_08
--==============================================================
create index cstn_subr_pcn_00_08 on cst_subr_pcn_00 (
	frst_pkgp_code ASC
);

--==============================================================
-- Index: cstn_subr_pcn_00_09
--==============================================================
create index cstn_subr_pcn_00_09 on cst_subr_pcn_00 (
	deal_numb ASC
);


--==============================================================
-- Index: cstn_subr_pcn_00_10
--==============================================================
create index cstn_subr_pcn_00_10 on cst_subr_pcn_00 (
	leas_strt_date, leas_code ASC
);

--==============================================================
-- Index: cstn_subr_pcn_00_11
--==============================================================
create index cstn_subr_pcn_00_11 on cst_subr_pcn_00 (
	frst_swon_dttm ASC
);

--==============================================================
-- Index: cstn_subr_pcn_00_12
--==============================================================
create index cstn_subr_pcn_00_12 on cst_subr_pcn_00 (
	brth_swon_dttm ASC
);

--==============================================================
-- Index: cstn_subr_pcn_00_13
--==============================================================
create index cstn_subr_pcn_00_13 on cst_subr_pcn_00 (
	subr_numb ASC
);