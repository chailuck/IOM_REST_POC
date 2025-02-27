create table CST_SPKD_EXT_00 (
	spkd_seqn integer  not null,
	cust_numb integer  not null,
	subr_numb char(12)  not null,
	pack_code char(8) not null,
	subr_row_id char(15) not null,
	bl_ext_id varchar(20)   ,
	ext_id varchar(70)   ,
	spkd_stts integer  ,
	rd_indc integer  ,
	offr_inst_id varchar(20)  ,
	parn_offr_inst_id varchar(20)  ,
	crtd_dttm datetime year to second  not null,
	crtd_by char(12)  not null,
	last_chng_dttm datetime year to second  not null,
	last_chng_by char(12)  not null,
	
	primary key (spkd_seqn)
		constraint cstp_spkd_ext_00_01
);


--==============================================================
-- Index: cstp_spkd_ext_00_01
--==============================================================
create unique index cstp_spkd_ext_00_01 on cst_spkd_ext_00 (
	spkd_seqn  ASC
);

--==============================================================
-- Index: cstn_spkd_ext_00_02
--==============================================================
create index cstn_spkd_ext_00_02 on cst_spkd_ext_00 (
	cust_numb, subr_numb, pack_code ASC
);


--==============================================================
-- Index: cstn_spkd_ext_00_03
--==============================================================
create index cstn_spkd_ext_00_03 on cst_spkd_ext_00 (
	spkd_stts ASC
);



--==============================================================
-- Index: cstn_spkd_ext_00_04
--==============================================================
create index cstn_spkd_ext_00_04 on cst_spkd_ext_00 (
	ext_id ASC
);


--==============================================================
-- Index: cstn_spkd_ext_00_05
--==============================================================
create index cstn_spkd_ext_00_05 on cst_spkd_ext_00 (
	bl_ext_id ASC
);


--==============================================================
-- Index: cstn_spkd_ext_00_06
--==============================================================
create index cstn_spkd_ext_00_06 on cst_spkd_ext_00 (
	subr_row_id, pack_code ASC
);
