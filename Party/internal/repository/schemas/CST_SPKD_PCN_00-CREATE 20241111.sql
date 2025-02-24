create table CST_SPKD_PCN_00 (
	seqn serial  not null,
	cust_numb integer  not null,
	subr_numb char(12)  not null,
	pkgp_code char(8),
	pack_code char(8)  not null,
	pack_type char(2), 
	artm_code smallint,
	disc_code smallint, 
	pack_strt_dttm datetime year to second  not null,
	pack_end_dttm datetime year to second  ,
	prov_strt_dttm datetime year to second  ,
	init_end_dttm datetime year to second  ,
	swof_dttm datetime year to second  ,
	expr_flag char(1)  not null,
	crtd_dttm datetime year to second  not null,
	crtd_by char(12)  not null,
	last_chng_dttm datetime year to second  not null,
	last_chng_by char(12)  not null,

	primary key (seqn)
		constraint cstp_spkd_pcn_00_01
);


--==============================================================
-- Index: cstp_spkd_pcn_00_01
--==============================================================
create unique index cstp_spkd_pcn_00_01 on cst_spkd_pcn_00 (
	seqn  ASC
);

--==============================================================
-- Index: cstp_spkd_pcn_00_02
--==============================================================
create index cstp_spkd_pcn_00_02 on cst_spkd_pcn_00 (
	cust_numb, subr_numb  ASC
);

--==============================================================
-- Index: cstp_spkd_pcn_00_03
--==============================================================
create index cstp_spkd_pcn_00_03 on cst_spkd_pcn_00 (
	pack_strt_dttm  ASC
);

--==============================================================
-- Index: cstp_spkd_pcn_00_04
--==============================================================
create index cstp_spkd_pcn_00_04 on cst_spkd_pcn_00 (
	pack_end_dttm  ASC
);


--==============================================================
-- Index: cstp_spkd_pcn_00_05
--==============================================================
create index cstp_spkd_pcn_00_05 on cst_spkd_pcn_00 (
	swof_dttm  ASC
);


--==============================================================
-- Index: cstp_spkd_pcn_00_06
--==============================================================
create index cstp_spkd_pcn_00_06 on cst_spkd_pcn_00 (
	expr_flag  ASC
);


--==============================================================
-- Index: cstp_spkd_pcn_00_07
--==============================================================
create index cstp_spkd_pcn_00_07 on cst_spkd_pcn_00 (
	pack_code  ASC
);
